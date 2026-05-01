package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 字典类型管理

func GetDictTypeList(c *gin.Context) {
	dictName := c.Query("dict_name")
	dictType := c.Query("dict_type")
	status := c.Query("status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.DictType{})

	if dictName != "" {
		query = query.Where("dict_name LIKE ?", "%"+dictName+"%")
	}
	if dictType != "" {
		query = query.Where("dict_type LIKE ?", "%"+dictType+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var dictTypes []models.DictType
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("id DESC").
		Find(&dictTypes)

	utils.Page(c, dictTypes, total, pageNum, pageSize)
}

func GetDictTypeById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var dictType models.DictType
	if err := config.DB.Where("id = ?", id).First(&dictType).Error; err != nil {
		utils.ErrorNotFound(c, "字典类型不存在")
		return
	}

	utils.Success(c, dictType)
}

func CreateDictType(c *gin.Context) {
	var dictType models.DictType
	if err := c.ShouldBindJSON(&dictType); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查字典类型是否存在
	var count int64
	config.DB.Model(&models.DictType{}).Where("dict_type = ?", dictType.DictType).Count(&count)
	if count > 0 {
		utils.ErrorBadRequest(c, "字典类型已存在")
		return
	}

	username, _ := c.Get("username")
	dictType.CreateBy = username.(string)

	if err := config.DB.Create(&dictType).Error; err != nil {
		utils.ErrorInternalServer(c, "创建字典类型失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", dictType)
}

func UpdateDictType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var dictType models.DictType
	if err := config.DB.Where("id = ?", id).First(&dictType).Error; err != nil {
		utils.ErrorNotFound(c, "字典类型不存在")
		return
	}

	var updateDictType models.DictType
	if err := c.ShouldBindJSON(&updateDictType); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查字典类型是否被其他类型使用
	if updateDictType.DictType != dictType.DictType {
		var count int64
		config.DB.Model(&models.DictType{}).Where("dict_type = ? AND id != ?", updateDictType.DictType, id).Count(&count)
		if count > 0 {
			utils.ErrorBadRequest(c, "字典类型已存在")
			return
		}
	}

	username, _ := c.Get("username")
	updateDictType.UpdateBy = username.(string)

	if err := config.DB.Model(&dictType).Updates(&updateDictType).Error; err != nil {
		utils.ErrorInternalServer(c, "更新字典类型失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteDictType(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.DictType{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除字典类型失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

// 字典数据管理

func GetDictDataList(c *gin.Context) {
	dictType := c.Query("dict_type")
	dictLabel := c.Query("dict_label")
	status := c.Query("status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.DictData{})

	if dictType != "" {
		query = query.Where("dict_type = ?", dictType)
	}
	if dictLabel != "" {
		query = query.Where("dict_label LIKE ?", "%"+dictLabel+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var dictDatas []models.DictData
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("dict_sort ASC, id ASC").
		Find(&dictDatas)

	utils.Page(c, dictDatas, total, pageNum, pageSize)
}

func GetDictDataById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var dictData models.DictData
	if err := config.DB.Where("id = ?", id).First(&dictData).Error; err != nil {
		utils.ErrorNotFound(c, "字典数据不存在")
		return
	}

	utils.Success(c, dictData)
}

func CreateDictData(c *gin.Context) {
	var dictData models.DictData
	if err := c.ShouldBindJSON(&dictData); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	dictData.CreateBy = username.(string)

	if err := config.DB.Create(&dictData).Error; err != nil {
		utils.ErrorInternalServer(c, "创建字典数据失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", dictData)
}

func UpdateDictData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var dictData models.DictData
	if err := config.DB.Where("id = ?", id).First(&dictData).Error; err != nil {
		utils.ErrorNotFound(c, "字典数据不存在")
		return
	}

	var updateDictData models.DictData
	if err := c.ShouldBindJSON(&updateDictData); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	updateDictData.UpdateBy = username.(string)

	if err := config.DB.Model(&dictData).Updates(&updateDictData).Error; err != nil {
		utils.ErrorInternalServer(c, "更新字典数据失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteDictData(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.DictData{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除字典数据失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
