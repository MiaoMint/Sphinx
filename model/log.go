package model

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Status  int    `json:"status" gorm:"index:idx_status"`
	Latency int    `json:"latency"`
	IP      string `json:"ip"`
	Method  string `json:"method" gorm:"index:idx_method"`
	Path    string `json:"path" gorm:"index:idx_path"`
	Domain  string `json:"domain" gorm:"index:idx_domain"`
	APIID   uint   `json:"api_id" gorm:"index"`
}
