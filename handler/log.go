package handler

import (
	"github.com/MiaoMint/Sphinx/database"
	"github.com/MiaoMint/Sphinx/model"
	"github.com/MiaoMint/Sphinx/result"
	"github.com/gofiber/fiber/v3"
)

func GetLogList(c fiber.Ctx) error {
	var logs []model.Log

	if err := database.DB.Order("id desc").Find(&logs).Error; err != nil {
		return err
	}

	return c.JSON(result.Success(logs))
}
