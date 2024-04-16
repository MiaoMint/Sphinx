package model

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Domain   string `json:"domain" gorm:"not null;unique;index"`
	APIs     []API  `json:"apis" gorm:"foreignKey:DomainID;constraint:OnDelete:CASCADE"`
	APICount int    `json:"api_count" gorm:"-"`
}
