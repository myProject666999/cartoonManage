package main

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	config.InitDB()

	// 自动迁移模型
	config.DB.AutoMigrate(
		&models.User{},
		&models.Department{},
		&models.Post{},
		&models.Menu{},
		&models.Role{},
		&models.DictType{},
		&models.DictData{},
		&models.Config{},
		&models.Notice{},
		&models.OperLog{},
		&models.LoginLog{},
		&models.File{},
		&models.Task{},
		&models.TaskLog{},
		&models.MallUser{},
		&models.Product{},
		&models.ProductCategory{},
		&models.ProductAttribute{},
		&models.Order{},
		&models.OrderItem{},
	)

	// 创建默认管理员
	createDefaultAdmin()

	// 初始化路由
	r := gin.Default()
	routers.InitRouter(r)

	// 启动服务器
	log.Printf("服务器启动在端口 %s", config.ServerConf.Port)
	if err := r.Run(":" + config.ServerConf.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

func createDefaultAdmin() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		// 创建默认部门
		dept := models.Department{
			DeptName: "总公司",
			OrderNum: 1,
			Status:   "0",
		}
		config.DB.Create(&dept)

		// 创建默认角色
		role := models.Role{
			RoleName: "超级管理员",
			RoleKey:  "admin",
			OrderNum: 1,
			Status:   "0",
		}
		config.DB.Create(&role)

		// 创建默认管理员用户
		admin := models.User{
			Username: "admin",
			Password: "admin123",
			Nickname: "管理员",
			Email:    "admin@example.com",
			Phone:    "13800138000",
			Gender:   "0",
			Status:   "0",
			DeptID:   dept.ID,
			Roles:    []models.Role{role},
		}
		config.DB.Create(&admin)

		log.Println("默认管理员账号已创建: admin / admin123")
	}
}
