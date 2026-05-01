package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context) {
	postCode := c.Query("post_code")
	postName := c.Query("post_name")
	status := c.Query("status")

	query := config.DB.Model(&models.Post{}).Where("deleted = 0")

	if postCode != "" {
		query = query.Where("post_code LIKE ?", "%"+postCode+"%")
	}
	if postName != "" {
		query = query.Where("post_name LIKE ?", "%"+postName+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var posts []models.Post
	query.Order("order_num ASC, id ASC").Find(&posts)

	utils.Success(c, posts)
}

func GetPostById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var post models.Post
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&post).Error; err != nil {
		utils.ErrorNotFound(c, "岗位不存在")
		return
	}

	utils.Success(c, post)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	post.CreateBy = username.(string)

	if err := config.DB.Create(&post).Error; err != nil {
		utils.ErrorInternalServer(c, "创建岗位失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", post)
}

func UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var post models.Post
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&post).Error; err != nil {
		utils.ErrorNotFound(c, "岗位不存在")
		return
	}

	var updatePost models.Post
	if err := c.ShouldBindJSON(&updatePost); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	updatePost.UpdateBy = username.(string)

	if err := config.DB.Model(&post).Updates(&updatePost).Error; err != nil {
		utils.ErrorInternalServer(c, "更新岗位失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeletePost(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	if err := config.DB.Model(&models.Post{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"deleted":   1,
			"update_by": username.(string),
		}).Error; err != nil {
		utils.ErrorInternalServer(c, "删除岗位失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

func GetMenuList(c *gin.Context) {
	menuName := c.Query("menu_name")
	status := c.Query("status")

	query := config.DB.Model(&models.Menu{})

	if menuName != "" {
		query = query.Where("menu_name LIKE ?", "%"+menuName+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var menus []models.Menu
	query.Order("order_num ASC, id ASC").Find(&menus)

	utils.Success(c, menus)
}

func GetMenuById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var menu models.Menu
	if err := config.DB.Where("id = ?", id).First(&menu).Error; err != nil {
		utils.ErrorNotFound(c, "菜单不存在")
		return
	}

	utils.Success(c, menu)
}

func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	menu.CreateBy = username.(string)

	if err := config.DB.Create(&menu).Error; err != nil {
		utils.ErrorInternalServer(c, "创建菜单失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", menu)
}

func UpdateMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var menu models.Menu
	if err := config.DB.Where("id = ?", id).First(&menu).Error; err != nil {
		utils.ErrorNotFound(c, "菜单不存在")
		return
	}

	var updateMenu models.Menu
	if err := c.ShouldBindJSON(&updateMenu); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	updateMenu.UpdateBy = username.(string)

	if err := config.DB.Model(&menu).Updates(&updateMenu).Error; err != nil {
		utils.ErrorInternalServer(c, "更新菜单失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var childCount int64
	config.DB.Model(&models.Menu{}).Where("parent_id = ?", id).Count(&childCount)
	if childCount > 0 {
		utils.ErrorBadRequest(c, "存在子菜单，不能删除")
		return
	}

	if err := config.DB.Delete(&models.Menu{}, id).Error; err != nil {
		utils.ErrorInternalServer(c, "删除菜单失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
