package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMallUserList(c *gin.Context) {
	username := c.Query("username")
	status := c.Query("status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.MallUser{}).Where("deleted = 0")

	if username != "" {
		query = query.Where("username LIKE ? OR nickname LIKE ?", "%"+username+"%", "%"+username+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var mallUsers []models.MallUser
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("id DESC").
		Find(&mallUsers)

	utils.Page(c, mallUsers, total, pageNum, pageSize)
}

func GetMallUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var mallUser models.MallUser
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&mallUser).Error; err != nil {
		utils.ErrorNotFound(c, "商城用户不存在")
		return
	}

	utils.Success(c, mallUser)
}

func UpdateMallUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var mallUser models.MallUser
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&mallUser).Error; err != nil {
		utils.ErrorNotFound(c, "商城用户不存在")
		return
	}

	var updateMallUser models.MallUser
	if err := c.ShouldBindJSON(&updateMallUser); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Model(&mallUser).Updates(&updateMallUser).Error; err != nil {
		utils.ErrorInternalServer(c, "更新商城用户失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteMallUser(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Model(&models.MallUser{}).
		Where("id IN ?", ids).
		Update("deleted", 1).Error; err != nil {
		utils.ErrorInternalServer(c, "删除商城用户失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

func GetProductList(c *gin.Context) {
	productName := c.Query("product_name")
	status := c.Query("status")
	categoryID := c.Query("category_id")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.Product{}).Where("deleted = 0")

	if productName != "" {
		query = query.Where("product_name LIKE ?", "%"+productName+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	var total int64
	query.Count(&total)

	var products []models.Product
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("id DESC").
		Find(&products)

	utils.Page(c, products, total, pageNum, pageSize)
}

func GetProductById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var product models.Product
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&product).Error; err != nil {
		utils.ErrorNotFound(c, "商品不存在")
		return
	}

	utils.Success(c, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		utils.ErrorInternalServer(c, "创建商品失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", product)
}

func UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var product models.Product
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&product).Error; err != nil {
		utils.ErrorNotFound(c, "商品不存在")
		return
	}

	var updateProduct models.Product
	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Model(&product).Updates(&updateProduct).Error; err != nil {
		utils.ErrorInternalServer(c, "更新商品失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteProduct(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Model(&models.Product{}).
		Where("id IN ?", ids).
		Update("deleted", 1).Error; err != nil {
		utils.ErrorInternalServer(c, "删除商品失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

func GetOrderList(c *gin.Context) {
	orderNo := c.Query("order_no")
	userName := c.Query("user_name")
	orderStatus := c.Query("order_status")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.Order{}).Where("deleted = 0")

	if orderNo != "" {
		query = query.Where("order_no LIKE ?", "%"+orderNo+"%")
	}
	if userName != "" {
		query = query.Where("user_name LIKE ?", "%"+userName+"%")
	}
	if orderStatus != "" {
		query = query.Where("order_status = ?", orderStatus)
	}

	var total int64
	query.Count(&total)

	var orders []models.Order
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("id DESC").
		Find(&orders)

	utils.Page(c, orders, total, pageNum, pageSize)
}

func GetOrderById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var order models.Order
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&order).Error; err != nil {
		utils.ErrorNotFound(c, "订单不存在")
		return
	}

	var orderItems []models.OrderItem
	config.DB.Where("order_id = ?", id).Find(&orderItems)
	order.OrderItems = orderItems

	utils.Success(c, order)
}

func UpdateOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var order models.Order
	if err := config.DB.Where("id = ? AND deleted = 0", id).First(&order).Error; err != nil {
		utils.ErrorNotFound(c, "订单不存在")
		return
	}

	var updateOrder models.Order
	if err := c.ShouldBindJSON(&updateOrder); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Model(&order).Updates(&updateOrder).Error; err != nil {
		utils.ErrorInternalServer(c, "更新订单失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func GetProductCategoryList(c *gin.Context) {
	var categories []models.ProductCategory
	config.DB.Order("order_num ASC, id ASC").Find(&categories)
	utils.Success(c, categories)
}

func GetProductCategoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var category models.ProductCategory
	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		utils.ErrorNotFound(c, "商品分类不存在")
		return
	}

	utils.Success(c, category)
}

func CreateProductCategory(c *gin.Context) {
	var category models.ProductCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		utils.ErrorInternalServer(c, "创建商品分类失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", category)
}

func UpdateProductCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var category models.ProductCategory
	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		utils.ErrorNotFound(c, "商品分类不存在")
		return
	}

	var updateCategory models.ProductCategory
	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Model(&category).Updates(&updateCategory).Error; err != nil {
		utils.ErrorInternalServer(c, "更新商品分类失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteProductCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var childCount int64
	config.DB.Model(&models.ProductCategory{}).Where("parent_id = ?", id).Count(&childCount)
	if childCount > 0 {
		utils.ErrorBadRequest(c, "存在子分类，不能删除")
		return
	}

	var productCount int64
	config.DB.Model(&models.Product{}).Where("category_id = ?", id).Count(&productCount)
	if productCount > 0 {
		utils.ErrorBadRequest(c, "分类下存在商品，不能删除")
		return
	}

	if err := config.DB.Delete(&models.ProductCategory{}, id).Error; err != nil {
		utils.ErrorInternalServer(c, "删除商品分类失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
