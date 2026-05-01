package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetFileList(c *gin.Context) {
	fileName := c.Query("file_name")
	fileType := c.Query("file_type")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.File{})

	if fileName != "" {
		query = query.Where("file_name LIKE ?", "%"+fileName+"%")
	}
	if fileType != "" {
		query = query.Where("file_type LIKE ?", "%"+fileType+"%")
	}

	var total int64
	query.Count(&total)

	var files []models.File
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("create_time DESC").
		Find(&files)

	utils.Page(c, files, total, pageNum, pageSize)
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		utils.ErrorBadRequest(c, "请选择上传文件")
		return
	}
	defer file.Close()

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	dateDir := filepath.Join(uploadDir, time.Now().Format("20060102"))
	if _, err := os.Stat(dateDir); os.IsNotExist(err) {
		os.MkdirAll(dateDir, 0755)
	}

	ext := filepath.Ext(header.Filename)
	uuidName := uuid.New().String() + ext
	savePath := filepath.Join(dateDir, uuidName)

	out, err := os.Create(savePath)
	if err != nil {
		utils.ErrorInternalServer(c, "创建文件失败")
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		utils.ErrorInternalServer(c, "保存文件失败")
		return
	}

	username, _ := c.Get("username")
	fileRecord := models.File{
		FileName:   header.Filename,
		FilePath:   strings.ReplaceAll(savePath, "\\", "/"),
		FileSize:   header.Size,
		FileType:   ext,
		MimeType:   header.Header.Get("Content-Type"),
		CreateBy:   username.(string),
		CreateTime: time.Now(),
	}

	if err := config.DB.Create(&fileRecord).Error; err != nil {
		utils.ErrorInternalServer(c, "保存文件记录失败")
		return
	}

	utils.SuccessWithMessage(c, "上传成功", gin.H{
		"id":       fileRecord.ID,
		"name":     fileRecord.FileName,
		"path":     fileRecord.FilePath,
		"size":     fileRecord.FileSize,
		"url":      fmt.Sprintf("/%s", strings.ReplaceAll(fileRecord.FilePath, "\\", "/")),
		"mimeType": fileRecord.MimeType,
	})
}

func DeleteFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var fileRecord models.File
	if err := config.DB.Where("id = ?", id).First(&fileRecord).Error; err != nil {
		utils.ErrorNotFound(c, "文件不存在")
		return
	}

	if _, err := os.Stat(fileRecord.FilePath); err == nil {
		os.Remove(fileRecord.FilePath)
	}

	if err := config.DB.Delete(&fileRecord).Error; err != nil {
		utils.ErrorInternalServer(c, "删除文件失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}

func ServeFile(c *gin.Context) {
	filePath := c.Param("filepath")
	fullPath := filepath.Join("./uploads", filePath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.Status(http.StatusNotFound)
		return
	}

	c.File(fullPath)
}
