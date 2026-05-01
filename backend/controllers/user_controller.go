package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Username string `form:"username"`
	Status   string `form:"status"`
	DeptID   uint64 `form:"dept_id"`
	PageNum  int    `form:"page_num,default=1"`
	PageSize int    `form:"page_size,default=10"`
}

func GetUserList(c *gin.Context) {
	var req UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	query := config.DB.Model(&models.User{}).Where("deleted = 0")

	if req.Username != "" {
		query = query.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}
	if req.DeptID > 0 {
		query = query.Where("dept_id = ?", req.DeptID)
	}

	var total int64
	query.Count(&total)

	var users []models.User
	offset := (req.PageNum - 1) * req.PageSize
	query.Preload("Roles").Preload("Posts").
		Offset(offset).Limit(req.PageSize).
		Order("id DESC").
		Find(&users)

	utils.Page(c, users, total, req.PageNum, req.PageSize)
}

func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var user models.User
	if err := config.DB.Preload("Roles").Preload("Posts").
		Where("id = ? AND deleted = 0", id).First(&user).Error; err != nil {
		utils.ErrorNotFound(c, "用户不存在")
		return
	}

	utils.Success(c, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查用户名是否存在
	var count int64
	config.DB.Model(&models.User{}).Where("username = ? AND deleted = 0", user.Username).Count(&count)
	if count > 0 {
		utils.ErrorBadRequest(c, "用户名已存在")
		return
	}

	// 设置创建者
	username, _ := c.Get("username")
	user.CreateBy = username.(string)

	if err := config.DB.Create(&user).Error; err != nil {
		utils.ErrorInternalServer(c, "创建用户失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", user)
}

func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var user models.User
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&user).Error; err != nil {
		utils.ErrorNotFound(c, "用户不存在")
		return
	}

	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查用户名是否被其他用户使用
	if updateUser.Username != user.Username {
		var count int64
		config.DB.Model(&models.User{}).Where("username = ? AND id != ? AND deleted = 0", updateUser.Username, id).Count(&count)
		if count > 0 {
			utils.ErrorBadRequest(c, "用户名已存在")
			return
		}
	}

	// 设置更新者
	username, _ := c.Get("username")
	updateUser.UpdateBy = username.(string)

	if err := config.DB.Model(&user).Updates(&updateUser).Error; err != nil {
		utils.ErrorInternalServer(c, "更新用户失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteUser(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 逻辑删除
	username, _ := c.Get("username")
	if err := config.DB.Model(&models.User{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"deleted":   1,
			"update_by": username.(string),
		}).Error; err != nil {
		utils.ErrorInternalServer(c, "删除用户失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

func ResetPassword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var req struct {
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.ErrorInternalServer(c, "密码加密失败")
		return
	}

	username, _ := c.Get("username")
	if err := config.DB.Model(&models.User{}).
		Where("id = ? AND deleted = 0", id).
		Updates(map[string]interface{}{
			"password":  hashedPassword,
			"update_by": username.(string),
		}).Error; err != nil {
		utils.ErrorInternalServer(c, "重置密码失败")
		return
	}

	utils.SuccessWithMessage(c, "重置密码成功", nil)
}

func UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	updateUser.UpdateBy = username.(string)

	if err := config.DB.Model(&models.User{}).
		Where("id = ? AND deleted = 0", userID).
		Updates(&updateUser).Error; err != nil {
		utils.ErrorInternalServer(c, "更新个人信息失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func UpdateCurrentPassword(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var user models.User
	if err := config.DB.Where("id = ? AND deleted = 0", userID).First(&user).Error; err != nil {
		utils.ErrorNotFound(c, "用户不存在")
		return
	}

	if !utils.CheckPassword(req.OldPassword, user.Password) {
		utils.ErrorBadRequest(c, "原密码错误")
		return
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		utils.ErrorInternalServer(c, "密码加密失败")
		return
	}

	username, _ := c.Get("username")
	if err := config.DB.Model(&user).
		Updates(map[string]interface{}{
			"password":  hashedPassword,
			"update_by": username.(string),
		}).Error; err != nil {
		utils.ErrorInternalServer(c, "更新密码失败")
		return
	}

	utils.SuccessWithMessage(c, "密码更新成功", nil)
}
