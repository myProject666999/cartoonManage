package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetOperLogList(c *gin.Context) {
	title := c.Query("title")
	operName := c.Query("oper_name")
	businessType := c.Query("business_type")
	status := c.Query("status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.OperLog{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if operName != "" {
		query = query.Where("oper_name LIKE ?", "%"+operName+"%")
	}
	if businessType != "" {
		query = query.Where("business_type = ?", businessType)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var operLogs []models.OperLog
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("oper_time DESC").
		Find(&operLogs)

	utils.Page(c, operLogs, total, pageNum, pageSize)
}

func GetOperLogById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var operLog models.OperLog
	if err := config.DB.Where("id = ?", id).First(&operLog).Error; err != nil {
		utils.ErrorNotFound(c, "操作日志不存在")
		return
	}

	utils.Success(c, operLog)
}

func DeleteOperLog(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.OperLog{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除操作日志失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

func GetLoginLogList(c *gin.Context) {
	loginName := c.Query("login_name")
	status := c.Query("status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.LoginLog{})

	if loginName != "" {
		query = query.Where("login_name LIKE ?", "%"+loginName+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var loginLogs []models.LoginLog
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("login_time DESC").
		Find(&loginLogs)

	utils.Page(c, loginLogs, total, pageNum, pageSize)
}

func DeleteLoginLog(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.LoginLog{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除登录日志失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
