package handler

import (
	"github.com/MiaoMint/Sphinx/database"
	"github.com/MiaoMint/Sphinx/host"
	"github.com/MiaoMint/Sphinx/model"
	"github.com/MiaoMint/Sphinx/result"
	"github.com/gofiber/fiber/v3"
)

func GetApiByDomain(c fiber.Ctx) error {
	id := c.Params("id")
	var apis []model.API
	if err := database.DB.Where("domain_id = ?", id).Find(&apis).Error; err != nil {
		return err
	}

	return c.JSON(result.Success(apis))
}

func PostApi(c fiber.Ctx) error {
	id := c.Params("id")

	domain := model.Domain{}
	if err := database.DB.First(&domain, id).Error; err != nil {
		return err
	}

	var api model.API
	if err := c.Bind().Body(&api); err != nil {
		return err
	}

	if err := database.DB.Model(&domain).Association("APIs").Append(&api); err != nil {
		return err
	}

	host.LoadHosts()

	return c.JSON(result.Success(api))
}

func PutApi(c fiber.Ctx) error {
	var api model.API

	err := c.Bind().Body(&api)
	if err != nil {
		return err
	}

	err = database.DB.Save(&api).Error
	if err != nil {
		return err
	}

	host.LoadHosts()
	return c.JSON(result.Success(api))
}

func DeleteApi(c fiber.Ctx) error {
	id := c.Params("id")
	err := database.DB.Unscoped().Delete(&model.API{}, id).Error
	if err != nil {
		return err
	}
	host.LoadHosts()
	return c.JSON(result.Success(nil))
}
