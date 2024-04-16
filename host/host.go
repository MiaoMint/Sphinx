package host

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/MiaoMint/Sphinx/cert"
	"github.com/MiaoMint/Sphinx/config"
	"github.com/MiaoMint/Sphinx/database"
	"github.com/MiaoMint/Sphinx/model"
	"github.com/MiaoMint/Sphinx/result"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

var (
	hosts = make(map[string]*fiber.App)
)

func LoadHosts() {
	var domains []model.Domain
	err := database.DB.Preload("APIs").Find(&domains).Error
	if err != nil {
		log.Fatalln(err)
	}
	for _, domain := range domains {
		AddHost(domain)
		// 判断是否需要生成证书
		_, certErr := os.Stat(cert.GetCertPath(domain.Domain))
		_, keyErr := os.Stat(cert.GetKeyPath(domain.Domain))
		if os.IsNotExist(certErr) || os.IsNotExist(keyErr) {
			if err := cert.GenerateClientCert(domain.Domain, config.LocalIP); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func AddHost(domain model.Domain) {
	config.AddHosts(domain.Domain)
	host := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(result.Fail(err.Error()))
		},
	})

	apisMap := make(map[string]map[string]model.API)

	for _, api := range domain.APIs {
		if _, ok := apisMap[api.Path]; !ok {
			apisMap[api.Path] = make(map[string]model.API)
		}
		apisMap[api.Path][api.Method] = api
	}

	host.All("/*", func(c fiber.Ctx) error {
		defer c.Next()
		c.Locals("time", time.Now())
		// 判断是否有对应的 api
		apis, ok := apisMap[c.Path()]
		// 判断请求方式
		api, methodOk := apis[c.Method()]

		if !ok || !methodOk {
			return proxy.DomainForward(domain.Domain, "https://"+domain.Domain, &fasthttp.Client{Dial: func(addr string) (net.Conn, error) {
				domain, port, err := net.SplitHostPort(addr)
				if err != nil {
					return nil, err
				}
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				// 通过 doh 查询 dns
				c := doh.Use(doh.DNSPodProvider)
				rsp, err := c.Query(ctx, dns.Domain(domain), dns.TypeA)
				if err != nil {
					return nil, err
				}
				c.Close()
				return net.Dial("tcp", net.JoinHostPort(rsp.Answer[0].Data, port))
			}})(c)
		}

		// 如有有对应的 api 则传递给下一个中间件
		c.Locals("APIID", api.ID)

		switch api.HandleMode {
		case model.HandleModeReplaceBody:
			return c.SendString(api.Body)
		case model.HandleModeModifyBody:
			return c.SendString(api.Body)
		case model.HandleModeJavaScript:
			return c.SendString(api.Body)
		default:
			return c.SendString(api.Body)
		}
	})

	host.Use(func(c fiber.Ctx) error {
		startTIme := c.Locals("time").(time.Time)
		apiID := c.Locals("APIID")
		if apiID == nil {
			return nil
		}

		logModel := model.Log{
			Status:  c.Response().StatusCode(),
			Latency: int(time.Since(startTIme).Milliseconds()),
			IP:      c.IP(),
			Method:  c.Method(),
			Path:    c.Path(),
			Domain:  domain.Domain,
		}

		err := database.DB.Model(&model.API{
			Model: gorm.Model{
				ID: apiID.(uint),
			},
		}).Association("Logs").Append(&logModel)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})

	hosts[domain.Domain] = host
}

func RemoveHost(domain string) {
	delete(hosts, domain)
}

func GetHost(domain string) *fiber.App {
	return hosts[domain]
}
