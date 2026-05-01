package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type ServerConfig struct {
	Port        string
	JWTSecret   string
	JWTExpire   time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var ServerConf *ServerConfig
var DatabaseConf *DatabaseConfig

func InitConfig() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Printf("警告: 未找到.env文件，使用默认配置")
	}

	// 初始化viper
	viper.AutomaticEnv()

	// 服务器配置
	ServerConf = &ServerConfig{
		Port:      getEnv("SERVER_PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "cartoon_manage_secret_key"),
		JWTExpire: time.Hour * 24, // 默认24小时过期
	}

	// 数据库配置
	DatabaseConf = &DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "root"),
		DBName:   getEnv("DB_NAME", "cartoon_manage"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func InitDB() {
	var err error
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseConf.User,
		DatabaseConf.Password,
		DatabaseConf.Host,
		DatabaseConf.Port,
		DatabaseConf.DBName,
	)

	DB, err = gorm.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 数据库连接池配置
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Hour)

	// 开启日志模式
	DB.LogMode(true)

	log.Println("数据库连接成功")
}
