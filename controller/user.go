package controller

import (
	"go-user-management/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// LoginReq 登录请求参数
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary 用户注册
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param body body controller.RegisterReq true "注册入参"
// @Success 200 {object} map[string]interface{}
// @Router /user/register [post]
func RegisterHandler(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数格式错误"})
		return
	}
	err := service.Register(req.Username, req.Password, req.Phone, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

// @Summary 用户登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param body body controller.LoginReq true "登录入参"
// @Success 200 {object} map[string]interface{}
// @Router /user/login [post]
func LoginHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}
	user, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"data": user,
	})
}
