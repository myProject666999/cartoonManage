package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetConfigList(c *gin.Context) {
	configName := c.Query("config_name")
	configType := c.Query("config_type")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.Config{})

	if configName != "" {
		query = query.Where("config_name LIKE ?", "%"+configName+"%")
	}
	if configType != "" {
		query = query.Where("config_type = ?", configType)
	}

	var total int64
	query.Count(&total)

	var configs []models.Config
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("id DESC").
		Find(&configs)

	utils.Page(c, configs, total, pageNum, pageSize)
}

func GetConfigById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var cfg models.Config
	if err := config.DB.Where("id = ?", id).First(&cfg).Error; err != nil {
		utils.ErrorNotFound(c, "参数不存在")
		return
	}

	utils.Success(c, cfg)
}

func CreateConfig(c *gin.Context) {
	var cfg models.Config
	if err := c.ShouldBindJSON(&cfg); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查参数键名是否存在
	var count int64
	config.DB.Model(&models.Config{}).Where("config_key = ?", cfg.ConfigKey).Count(&count)
	if count > 0 {
		utils.ErrorBadRequest(c, "参数键名已存在")
		return
	}

	username, _ := c.Get("username")
	cfg.CreateBy = username.(string)

	if err := config.DB.Create(&cfg).Error; err != nil {
		utils.ErrorInternalServer(c, "创建参数失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", cfg)
}

func UpdateConfig(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var cfg models.Config
	if err := config.DB.Where("id = ?", id).First(&cfg).Error; err != nil {
		utils.ErrorNotFound(c, "参数不存在")
		return
	}

	var updateConfig models.Config
	if err := c.ShouldBindJSON(&updateConfig); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查参数键名是否被其他参数使用
	if updateConfig.ConfigKey != cfg.ConfigKey {
		var count int64
		config.DB.Model(&models.Config{}).Where("config_key = ? AND id != ?", updateConfig.ConfigKey, id).Count(&count)
		if count > 0 {
			utils.ErrorBadRequest(c, "参数键名已存在")
			return
		}
	}

	username, _ := c.Get("username")
	updateConfig.UpdateBy = username.(string)

	if err := config.DB.Model(&cfg).Updates(&updateConfig).Error; err != nil {
		utils.ErrorInternalServer(c, "更新参数失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteConfig(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.Config{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除参数失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
