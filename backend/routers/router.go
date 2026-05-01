package routers

import (
	"cartoonManage/controllers"
	"cartoonManage/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// 允许跨域
	r.Use(CORSMiddleware())

	// 登录接口，不需要认证
	r.POST("/api/auth/login", controllers.Login)
	r.POST("/api/auth/logout", controllers.Logout)

	// 需要认证的接口
	auth := r.Group("/api")
	auth.Use(middlewares.JWTAuth())
	{
		// 获取用户信息
		auth.GET("/auth/info", controllers.GetUserInfo)
		auth.PUT("/auth/profile", controllers.UpdateProfile)
		auth.PUT("/auth/password", controllers.UpdateCurrentPassword)

		// 用户管理
		user := auth.Group("/system/user")
		{
			user.GET("/list", controllers.GetUserList)
			user.GET("/:id", controllers.GetUserById)
			user.POST("", controllers.CreateUser)
			user.PUT("/:id", controllers.UpdateUser)
			user.DELETE("/:ids", controllers.DeleteUser)
			user.PUT("/resetPwd/:id", controllers.ResetPassword)
		}

		// 部门管理
		dept := auth.Group("/system/dept")
		{
			dept.GET("/list", controllers.GetDeptList)
			dept.GET("/:id", controllers.GetDeptById)
			dept.POST("", controllers.CreateDept)
			dept.PUT("/:id", controllers.UpdateDept)
			dept.DELETE("/:id", controllers.DeleteDept)
		}

		// 岗位管理
		post := auth.Group("/system/post")
		{
			post.GET("/list", controllers.GetPostList)
			post.GET("/:id", controllers.GetPostById)
			post.POST("", controllers.CreatePost)
			post.PUT("/:id", controllers.UpdatePost)
			post.DELETE("/:ids", controllers.DeletePost)
		}

		// 菜单管理
		menu := auth.Group("/system/menu")
		{
			menu.GET("/list", controllers.GetMenuList)
			menu.GET("/:id", controllers.GetMenuById)
			menu.POST("", controllers.CreateMenu)
			menu.PUT("/:id", controllers.UpdateMenu)
			menu.DELETE("/:id", controllers.DeleteMenu)
		}

		// 角色管理
		role := auth.Group("/system/role")
		{
			role.GET("/list", controllers.GetRoleList)
			role.GET("/:id", controllers.GetRoleById)
			role.POST("", controllers.CreateRole)
			role.PUT("/:id", controllers.UpdateRole)
			role.DELETE("/:ids", controllers.DeleteRole)
		}

		// 字典类型管理
		dictType := auth.Group("/system/dict/type")
		{
			dictType.GET("/list", controllers.GetDictTypeList)
			dictType.GET("/:id", controllers.GetDictTypeById)
			dictType.POST("", controllers.CreateDictType)
			dictType.PUT("/:id", controllers.UpdateDictType)
			dictType.DELETE("/:ids", controllers.DeleteDictType)
		}

		// 字典数据管理
		dictData := auth.Group("/system/dict/data")
		{
			dictData.GET("/list", controllers.GetDictDataList)
			dictData.GET("/:id", controllers.GetDictDataById)
			dictData.POST("", controllers.CreateDictData)
			dictData.PUT("/:id", controllers.UpdateDictData)
			dictData.DELETE("/:ids", controllers.DeleteDictData)
		}

		// 参数管理
		config := auth.Group("/system/config")
		{
			config.GET("/list", controllers.GetConfigList)
			config.GET("/:id", controllers.GetConfigById)
			config.POST("", controllers.CreateConfig)
			config.PUT("/:id", controllers.UpdateConfig)
			config.DELETE("/:ids", controllers.DeleteConfig)
		}

		// 通知公告
		notice := auth.Group("/system/notice")
		{
			notice.GET("/list", controllers.GetNoticeList)
			notice.GET("/:id", controllers.GetNoticeById)
			notice.POST("", controllers.CreateNotice)
			notice.PUT("/:id", controllers.UpdateNotice)
			notice.DELETE("/:ids", controllers.DeleteNotice)
		}

		// 操作日志
		operLog := auth.Group("/monitor/operlog")
		{
			operLog.GET("/list", controllers.GetOperLogList)
			operLog.GET("/:id", controllers.GetOperLogById)
			operLog.DELETE("/:ids", controllers.DeleteOperLog)
		}

		// 登录日志
		loginLog := auth.Group("/monitor/logininfor")
		{
			loginLog.GET("/list", controllers.GetLoginLogList)
			loginLog.DELETE("/:ids", controllers.DeleteLoginLog)
		}

		// 文件管理
		file := auth.Group("/system/file")
		{
			file.GET("/list", controllers.GetFileList)
			file.POST("/upload", controllers.UploadFile)
			file.DELETE("/:id", controllers.DeleteFile)
		}

		// 在线用户
		online := auth.Group("/monitor/online")
		{
			online.GET("/list", controllers.GetOnlineUserList)
		}

		// 定时任务
		task := auth.Group("/monitor/job")
		{
			task.GET("/list", controllers.GetTaskList)
			task.GET("/:id", controllers.GetTaskById)
			task.POST("", controllers.CreateTask)
			task.PUT("/:id", controllers.UpdateTask)
			task.DELETE("/:ids", controllers.DeleteTask)
		}

		// 定时任务日志
		taskLog := auth.Group("/monitor/jobLog")
		{
			taskLog.GET("/list", controllers.GetTaskLogList)
			taskLog.DELETE("/:ids", controllers.DeleteTaskLog)
		}

		// 服务监控
		auth.GET("/monitor/server", controllers.GetServerInfo)
		auth.GET("/monitor/dbpool", controllers.GetDBPoolInfo)

		// 商城用户管理
		mallUser := auth.Group("/mall/user")
		{
			mallUser.GET("/list", controllers.GetMallUserList)
			mallUser.GET("/:id", controllers.GetMallUserById)
			mallUser.PUT("/:id", controllers.UpdateMallUser)
			mallUser.DELETE("/:ids", controllers.DeleteMallUser)
		}

		// 商品分类管理
		productCategory := auth.Group("/mall/product/category")
		{
			productCategory.GET("/list", controllers.GetProductCategoryList)
			productCategory.GET("/:id", controllers.GetProductCategoryById)
			productCategory.POST("", controllers.CreateProductCategory)
			productCategory.PUT("/:id", controllers.UpdateProductCategory)
			productCategory.DELETE("/:id", controllers.DeleteProductCategory)
		}

		// 商品管理
		product := auth.Group("/mall/product")
		{
			product.GET("/list", controllers.GetProductList)
			product.GET("/:id", controllers.GetProductById)
			product.POST("", controllers.CreateProduct)
			product.PUT("/:id", controllers.UpdateProduct)
			product.DELETE("/:ids", controllers.DeleteProduct)
		}

		// 订单管理
		order := auth.Group("/mall/order")
		{
			order.GET("/list", controllers.GetOrderList)
			order.GET("/:id", controllers.GetOrderById)
			order.PUT("/:id", controllers.UpdateOrder)
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
