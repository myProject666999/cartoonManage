package controllers

import (
	"cartoonManage/config"
	"cartoonManage/models"
	"cartoonManage/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetNoticeList(c *gin.Context) {
	noticeTitle := c.Query("notice_title")
	noticeType := c.Query("notice_type")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	query := config.DB.Model(&models.Notice{})

	if noticeTitle != "" {
		query = query.Where("notice_title LIKE ?", "%"+noticeTitle+"%")
	}
	if noticeType != "" {
		query = query.Where("notice_type = ?", noticeType)
	}

	var total int64
	query.Count(&total)

	var notices []models.Notice
	offset := (pageNum - 1) * pageSize
	query.Offset(offset).Limit(pageSize).
		Order("id DESC").
		Find(&notices)

	utils.Page(c, notices, total, pageNum, pageSize)
}

func GetNoticeById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var notice models.Notice
	if err := config.DB.Where("id = ?", id).First(&notice).Error; err != nil {
		utils.ErrorNotFound(c, "通知公告不存在")
		return
	}

	utils.Success(c, notice)
}

func CreateNotice(c *gin.Context) {
	var notice models.Notice
	if err := c.ShouldBindJSON(&notice); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	notice.CreateBy = username.(string)

	if err := config.DB.Create(&notice).Error; err != nil {
		utils.ErrorInternalServer(c, "创建通知公告失败")
		return
	}

	utils.SuccessWithMessage(c, "创建成功", notice)
}

func UpdateNotice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	var notice models.Notice
	if err := config.DB.Where("id = ?", id).First(&notice).Error; err != nil {
		utils.ErrorNotFound(c, "通知公告不存在")
		return
	}

	var updateNotice models.Notice
	if err := c.ShouldBindJSON(&updateNotice); err != nil {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	username, _ := c.Get("username")
	updateNotice.UpdateBy = username.(string)

	if err := config.DB.Model(&notice).Updates(&updateNotice).Error; err != nil {
		utils.ErrorInternalServer(c, "更新通知公告失败")
		return
	}

	utils.SuccessWithMessage(c, "更新成功", nil)
}

func DeleteNotice(c *gin.Context) {
	idsStr := c.Param("ids")
	ids := strings.Split(idsStr, ",")

	if len(ids) == 0 {
		utils.ErrorBadRequest(c, "参数错误")
		return
	}

	if err := config.DB.Delete(&models.Notice{}, ids).Error; err != nil {
		utils.ErrorInternalServer(c, "删除通知公告失败")
		return
	}

	utils.SuccessWithMessage(c, "删除成功", nil)
}
