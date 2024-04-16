package model

import "gorm.io/gorm"

type HandleMode string

var (
	HandleModeReplaceBody HandleMode = "ReplaceBody"
	HandleModeModifyBody  HandleMode = "ModifyBody"
	HandleModeJavaScript  HandleMode = "JavaScript"
)

type API struct {
	gorm.Model
	Name        string     `json:"name"`
	Path        string     `json:"path"`
	Method      string     `json:"method" gorm:"not null;default:'GET'"`
	HandleMode  HandleMode `json:"handle_mode" gorm:"not null"`
	Body        string     `json:"body"`
	JavaScript  string     `json:"javascript"`
	Replace     string     `json:"replace"`
	ReplaceWith string     `json:"replace_with"`
	DomainID    uint       `json:"domain_id" gorm:"index"`
	Logs        []Log      `json:"logs" gorm:"foreignKey:APIID;"`
}
