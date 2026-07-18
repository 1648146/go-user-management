package dao

import (
	"go-user-management/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(d *gorm.DB) {
	db = d
	// 自动创建用户表
	_ = db.AutoMigrate(&model.User{})
}

// CreateUser 创建用户
func CreateUser(user *model.User) error {
	return db.Create(user).Error
}

// GetUserByName 根据用户名查询
func GetUserByName(username string) (*model.User, error) {
	var u model.User
	err := db.Where("username = ?", username).First(&u).Error
	return &u, err
}

// ListUser 分页查询用户
func ListUser(page, size int) ([]model.User, int64, error) {
	var list []model.User
	var total int64
	db.Model(&model.User{}).Count(&total)
	err := db.Limit(size).Offset((page - 1) * size).Find(&list).Error
	return list, total, err
}

// UpdateUserStatus 修改账号状态
func UpdateUserStatus(uid uint, status int) error {
	return db.Model(&model.User{}).Where("id = ?", uid).Update("status", status).Error
}
