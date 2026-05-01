package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleListRequest struct {
	RoleName string `form:"role_name"`
	Status   string `form:"status"`
	PageNum  int    `form:"page_num,default=1"`
	PageSize int    `form:"page_size,default=10"`
}

func GetRoleList(c *gin.Context) {
	var req RoleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	query := config.DB.Model(&models.Role{}).Where("deleted = 0")

	if req.RoleName != "" {
		query = query.Where("role_name LIKE ?", "%"+req.RoleName+"%")
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	var total int64
	query.Count(&total)

	var roles []models.Role
	offset := (req.PageNum - 1) * req.PageSize
	query.Preload("Menus").
		Offset(offset).Limit(req.PageSize).
		Order("order_num ASC, id ASC").
		Find(&roles)

	utils.Page(c, roles, total, req.PageNum, req.PageSize)
}

func GetRoleById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var role models.Role
	if err := config.DB.Preload("Menus").
		Where("id = ? AND deleted = 0", id).First(&role).Error; err != nil {
		utils.ErrorNotFound(c, "角色不存在")
		return
	}

	utils.Success(c, role)
}

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查角色标识是否存在
	var count int64
	config.DB.Model(&models.Role{}).Where("role_key = ? AND deleted = 0", role.RoleKey).Count(&count)
	if count > 0 {
		utils.ErrorBadRequest(c, "角色标识已存在")
		return
	}

	username, _ := c.Get("username")
	role.CreateBy = username.(string)

	if err := config.DB.Create(&role).Error; err != nil {
		utils.ErrorInternalServer(c, "创建角色失败")
		return
	}

	// 关联菜单
	if len(role.MenuIDs) > 0 {
		var menus []models.Menu
		config.DB.Where("id IN ?", role.MenuIDs).Find(&menus)
		config.DB.Model(&role).Association("Menus").Replace(&menus)
	}

	utils.SuccessWithMessage(c, "创建成功", role)
}

func UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var role models.Role
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&role).Error; err != nil {
		utils.ErrorNotFound(c, "角色不存在")
		return
	}

	var updateRole models.Role
	if err := c.ShouldBindJSON(&updateRole); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查角色标识是否被其他角色使用
	if updateRole.RoleKey != role.RoleKey {
		var count int64
		config.DB.Model(&models.Role{}).Where("role_key = ? AND id != ? AND deleted = 0", updateRole.RoleKey, id).Count(&count)
		if count > 0 {
			utils.ErrorBadRequest(c, "角色标识已存在")
			return
		}
	}

	username, _ := c.Get("username")
	updateRole.UpdateBy = username.(string)

	if err := config.DB.Model(&role).Updates(&updateRole).Error; err != nil {
		utils.ErrorInternalServer(c, "更新角色失败")
		return
	}

	// 更新菜单关联
	if len(updateRole.MenuIDs) > 0 {
		var menus []models.Menu
		config.DB.Where("id IN ?", updateRole.MenuIDs).Find(&menus)
		config.DB.Model(&role).Association("Menus").Replace(&menus)
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteRole(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	// 检查角色是否有关联用户
	for _, idStr := range ids {
		id, _ := strconv.ParseUint(idStr, 10, 64)
		var count int64
		config.DB.Table("sys_user_role").Where("role_id = ?", id).Count(&count)
		if count > 0 {
			utils.ErrorBadRequest(c, "角色下存在用户，不能删除")
			return
		}
	}

	username, _ := c.Get("username")
	if err := config.DB.Model(&models.Role{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"deleted":   1,
			"update_by": username.(string),
		}).Error; err != nil {
		utils.ErrorInternalServer(c, "删除角色失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
