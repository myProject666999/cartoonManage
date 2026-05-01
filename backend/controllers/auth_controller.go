package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  *models.User `json:"user"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var user models.User
	if err := config.DB.Where("username = ? AND deleted = 0", req.Username).First(&user).Error; err != nil {
		utils.ErrorUnauthorized(c, "用户名或密码错误")
		return
	}

	if user.Status == "1" {
		utils.ErrorForbidden(c, "账号已被禁用")
		return
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		utils.ErrorUnauthorized(c, "用户名或密码错误")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Nickname)
	if err != nil {
		utils.ErrorInternalServer(c, "生成token失败")
		return
	}

	utils.Success(c, LoginResponse{
		Token: token,
		User:  &user,
	})
}

func GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorUnauthorized(c, "用户未登录")
		return
	}

	var user models.User
	if err := config.DB.Preload("Roles").Preload("Posts").
		Where("id = ? AND deleted = 0", userID).First(&user).Error; err != nil {
		utils.ErrorNotFound(c, "用户不存在")
		return
	}

	utils.Success(c, user)
}

func Logout(c *gin.Context) {
	utils.SuccessWithMessage(c, "退出登录成功", nil)
}
