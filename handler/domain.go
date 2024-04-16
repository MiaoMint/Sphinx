package handler

import (
	"github.com/MiaoMint/Sphinx/cert"
	"github.com/MiaoMint/Sphinx/config"
	"github.com/MiaoMint/Sphinx/database"
	"github.com/MiaoMint/Sphinx/host"
	"github.com/MiaoMint/Sphinx/model"
	"github.com/MiaoMint/Sphinx/result"
	"github.com/gofiber/fiber/v3"
)

func GetHostsFile(c fiber.Ctx) error {
	return c.JSON(result.Success(config.GetHosts()))
}

func GetDomainList(c fiber.Ctx) error {
	var domains []*model.Domain

	err := database.DB.Find(&domains).Error
	if err != nil {
		return err
	}

	for _, domain := range domains {
		domain.APICount = int(database.DB.Model(&domain).Association("APIs").Count())
	}

	return c.JSON(result.Success(domains))
}

func GetDomain(c fiber.Ctx) error {
	id := c.Params("id")

	var domain model.Domain
	err := database.DB.Where("id = ?", id).First(&domain).Error
	if err != nil {
		return err
	}

	return c.JSON(result.Success(domain))
}

func PostDomain(c fiber.Ctx) error {
	var domain model.Domain

	err := c.Bind().Body(&domain)
	if err != nil {
		return err
	}

	err = database.DB.Create(&domain).Error
	if err != nil {
		return err
	}

	host.AddHost(domain)
	err = cert.GenerateClientCert(domain.Domain, config.LocalIP)
	if err != nil {
		return err
	}
	config.AddHosts(domain.Domain)
	cert.RefreshCerts()

	return c.JSON(result.Success(domain))
}

func PutDomain(c fiber.Ctx) error {
	id := c.Params("id")
	var domain model.Domain

	err := c.Bind().Body(&domain)
	if err != nil {
		return err
	}

	err = database.DB.Model(&model.Domain{}).Where("id = ?", id).Updates(&domain).Error
	if err != nil {
		return err
	}

	return c.JSON(result.Success(domain))
}

func DeleteDomain(c fiber.Ctx) error {
	id := c.Params("id")

	var domain model.Domain
	err := database.DB.Where("id = ?", id).First(&domain).Error
	if err != nil {
		return err
	}

	// 删除 apis
	if err := database.DB.Where("domain_id = ?", id).Unscoped().Delete(&model.API{}).Error; err != nil {
		return err
	}

	err = database.DB.Unscoped().Delete(&model.Domain{}, id).Error
	if err != nil {
		return err
	}

	config.RemoveHosts(domain.Domain)
	cert.RemoveCert(domain.Domain)
	cert.RefreshCerts()
	host.RemoveHost(domain.Domain)

	return c.JSON(result.Success(nil))
}
