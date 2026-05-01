package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDeptList(c *gin.Context) {
	deptName := c.Query("dept_name")
	status := c.Query("status")

	query := config.DB.Model(&models.Department{}).Where("deleted = 0")

	if deptName != "" {
		query = query.Where("dept_name LIKE ?", "%"+deptName+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var depts []models.Department
	query.Order("order_num ASC, id ASC").Find(&depts)

	utils.Success(c, depts)
}

func GetDeptById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var dept models.Department
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&dept).Error; err != nil {
		utils.ErrorNotFound(c, "部门不存在")
		return
	}

	utils.Success(c, dept)
}

func CreateDept(c *gin.Context) {
	var dept models.Department
	if err := c.ShouldBindJSON(&dept); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	dept.CreateBy = username.(string)

	if err := config.DB.Create(&dept).Error; err != nil {
		utils.ErrorInternalServer(c, "创建部门失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", dept)
}

func UpdateDept(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var dept models.Department
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&dept).Error; err != nil {
		utils.ErrorNotFound(c, "部门不存在")
		return
	}

	var updateDept models.Department
	if err := c.ShouldBindJSON(&updateDept); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	updateDept.UpdateBy = username.(string)

	if err := config.DB.Model(&dept).Updates(&updateDept).Error; err != nil {
		utils.ErrorInternalServer(c, "更新部门失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteDept(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查是否有子部门
	var childCount int64
	config.DB.Model(&models.Department{}).Where("parent_id = ? AND deleted = 0", id).Count(&childCount)
	if childCount > 0 {
		utils.ErrorBadRequest(c, "存在子部门，不能删除")
		return
	}

	// 检查是否有用户
	var userCount int64
	config.DB.Model(&models.User{}).Where("dept_id = ? AND deleted = 0", id).Count(&userCount)
	if userCount > 0 {
		utils.ErrorBadRequest(c, "部门下存在用户，不能删除")
		return
	}

	// 逻辑删除
	username, _ := c.Get("username")
	if err := config.DB.Model(&models.Department{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":   1,
			"update_by": username.(string),
		}).Error; err != nil {
		utils.ErrorInternalServer(c, "删除部门失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
