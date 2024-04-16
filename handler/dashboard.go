package handler

import (
	"time"

	"github.com/MiaoMint/Sphinx/database"
	"github.com/MiaoMint/Sphinx/model"
	"github.com/MiaoMint/Sphinx/result"
	"github.com/gofiber/fiber/v3"
)

func GetTotal(c fiber.Ctx) error {
	data := make(map[string]interface{})

	// 统计域名数量
	var domainCount int64
	if err := database.DB.Model(&model.Domain{}).Count(&domainCount).Error; err != nil {
		return err
	}

	// 统计 API 数量
	var apiCount int64
	if err := database.DB.Model(&model.API{}).Count(&apiCount).Error; err != nil {
		return err
	}

	// 统计总请求数量
	var requestCount int64
	if err := database.DB.Model(&model.Log{}).Count(&requestCount).Error; err != nil {
		return err
	}

	// 统计成功请求数量
	var successCount int64
	if err := database.DB.Model(&model.Log{}).Where("status = ?", 200).Count(&successCount).Error; err != nil {
		return err
	}

	data["domain_count"] = domainCount
	data["api_count"] = apiCount
	data["request_count"] = requestCount
	data["success_count"] = successCount

	return c.JSON(result.Success(data))
}

func GetOverview(c fiber.Ctx) error {
	// 统计最近 7 天的每天的请求数
	startTime := time.Now().AddDate(0, 0, -6)
	endTime := time.Now()

	var logs []model.Log
	if err := database.DB.Where("created_at BETWEEN ? AND ?", startTime, endTime).Find(&logs).Error; err != nil {
		return err
	}

	var data []map[string]interface{}
	for i := 0; i < 7; i++ {
		date := startTime.AddDate(0, 0, i).Format("2006-01-02")
		count := 0
		for _, log := range logs {
			if log.CreatedAt.Format("2006-01-02") == date {
				count++
			}
		}
		data = append(data, map[string]interface{}{
			"date":  date,
			"count": count,
			"index": i,
		})
	}
	return c.JSON(result.Success(data))
}

// 获取每个域名的请求排行
func GetDomainRank(c fiber.Ctx) error {
	var data []struct {
		Domain string `json:"domain"`
		Count  int    `json:"count"`
		Index  int    `json:"index"`
	}

	if err := database.DB.Model(&model.Log{}).Select("domain, count(*) as count").Group("domain").Order("count desc").Limit(10).Find(&data).Error; err != nil {
		return err
	}

	return c.JSON(result.Success(data))
}

// 获取每个 API 的请求排行
func GetAPIRank(c fiber.Ctx) error {
	var data []struct {
		Path   string `json:"path"`
		Domain string `json:"domain"`
		Method string `json:"method"`
		Count  int    `json:"count"`
		Index  int    `json:"index"`
	}

	if err := database.DB.Model(&model.Log{}).
		Select("api_id, count(*) as count, path,method,domain").
		Group("api_id").
		Order("count desc").
		Limit(10).
		Find(&data).
		Error; err != nil {
		return err
	}

	return c.JSON(result.Success(data))
}
