package models

import (
	"comment/models/stype"
)

type UserModels struct {
	MODEL
	Username string `gorm:"size:32;not null;unique" json:"user_name"`
	NickName string `gorm:"size=32" json:"nick_name"`
	Password string `gorm:"size:64;not null" json:"password"`
	Token    string `gorm:"64" json:"token"`
	IP       string `gorm:"size:20" json:"ip"`
	// admin:1 user:2 normal:3 banned:4
	Permission stype.Permission `gorm:"size:4;not null;default:1" json:"permission"`
	AccessKey  string           `json:"access_key"`
	SecretKey  string           `json:"secret_key"`

	SignStatus stype.SignStatus `gorm:"type=smallint(6);not null" json:"sign_status"`
}
