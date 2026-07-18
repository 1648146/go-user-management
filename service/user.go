package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go-user-management/dao"
	"go-user-management/model"
)

// Md5Encrypt 密码加密
func Md5Encrypt(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Register 用户注册
func Register(username, password, phone, email string) error {
	_, err := dao.GetUserByName(username)
	if err == nil {
		return errors.New("用户名已被注册")
	}
	user := model.User{
		Username: username,
		Password: Md5Encrypt(password),
		Phone:    phone,
		Email:    email,
	}
	return dao.CreateUser(&user)
}

// Login 用户登录校验
func Login(username, password string) (*model.User, error) {
	user, err := dao.GetUserByName(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if user.Status == 0 {
		return nil, errors.New("账号已被禁用")
	}
	if user.Password != Md5Encrypt(password) {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
