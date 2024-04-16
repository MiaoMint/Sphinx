package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/MiaoMint/Sphinx/cert"
	"github.com/MiaoMint/Sphinx/config"
	_ "github.com/MiaoMint/Sphinx/database"
	"github.com/MiaoMint/Sphinx/handler"
	"github.com/MiaoMint/Sphinx/host"
	"github.com/MiaoMint/Sphinx/result"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func main() {
	host.LoadHosts()
	server := fiber.New(fiber.Config{
		AppName: "Sphinx",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(result.Fail(err.Error()))
		},
	})

	server.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${domain} | ${method} | ${path} ${error}\n",
		CustomTags: map[string]logger.LogFunc{
			"domain": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(c.Hostname())
			},
		},
	}))

	server.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))

	server.Use(func(c fiber.Ctx) error {
		host := host.GetHost(c.Hostname())
		if host == nil {
			return c.Next()
		} else {
			host.Handler()(c.Context())
			return nil
		}
	})

	api := server.Group("/api")
	api.Get("/dashboard/total", handler.GetTotal)
	api.Get("/dashboard/overview", handler.GetOverview)
	api.Get("/dashboard/domain", handler.GetDomainRank)
	api.Get("/dashboard/api", handler.GetAPIRank)
	api.Get("/log", handler.GetLogList)

	api.Get("/hosts_file", handler.GetHostsFile)
	api.Get("/domain", handler.GetDomainList)
	api.Get("/domain/:id", handler.GetDomain)
	api.Post("/domain", handler.PostDomain)
	api.Put("/domain/:id", handler.PutDomain)
	api.Delete("/domain/:id", handler.DeleteDomain)

	api.Get("/domain/:id/api", handler.GetApiByDomain)
	api.Post("/domain/:id/api", handler.PostApi)
	api.Put("/domain/:domain_id/api/:id", handler.PutApi)
	api.Delete("/domain/:domain_id/api/:id", handler.DeleteApi)

	if config.DevMode {
		server.All("/*", proxy.DomainForward(config.Domain, "http://localhost:3000"))
	} else {
		server.Static("/", "./frontend/.output/public")
	}

	if err := cert.GenerateClientCert(config.Domain, config.LocalIP); err != nil {
		log.Fatalln(err)
	}

	cert.RefreshCerts()

	ln, err := tls.Listen("tcp", ":443", cert.TlsConfig)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(`
   _____       _     _            
  / ____|     | |   (_)           
 | (___  _ __ | |__  _ _ __ __  __
  \___ \| '_ \| '_ \| | '_ \\ \/ /
  ____) | |_) | | | | | | | |>  < 
 |_____/| .__/|_| |_|_|_| |_/_/\_\
        | |                       
        |_|                       
	`)
	fmt.Printf("Sphinx is running on https://%s\n", config.Domain)

	go func() {
		err := server.Listen(":80", fiber.ListenConfig{
			DisableStartupMessage: true,
		})
		if err != nil {
			log.Fatalln(err)
		}
	}()

	err = server.Listener(ln, fiber.ListenConfig{
		DisableStartupMessage: true,
	})

	if err != nil {
		log.Fatalln(err)
	}
}
