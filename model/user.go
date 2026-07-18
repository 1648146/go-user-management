package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:32;unique;comment:用户名"`
	Password string `gorm:"size:64;comment:加密密码"`
	Phone    string `gorm:"size:11;comment:手机号"`
	Email    string `gorm:"size:64;comment:邮箱"`
	Status   int    `gorm:"default:1;comment:1正常 0禁用"`
}
