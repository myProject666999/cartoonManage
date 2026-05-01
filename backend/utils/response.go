package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PageResponse struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "操作成功",
		Data:    data,
	})
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

func ErrorBadRequest(c *gin.Context, message string) {
	Error(c, 400, message)
}

func ErrorUnauthorized(c *gin.Context, message string) {
	Error(c, 401, message)
}

func ErrorForbidden(c *gin.Context, message string) {
	Error(c, 403, message)
}

func ErrorNotFound(c *gin.Context, message string) {
	Error(c, 404, message)
}

func ErrorInternalServer(c *gin.Context, message string) {
	Error(c, 500, message)
}

func Page(c *gin.Context, list interface{}, total int64, pageNum, pageSize int) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "操作成功",
		Data: PageResponse{
			List:     list,
			Total:    total,
			PageNum:  pageNum,
			PageSize: pageSize,
		},
	})
}
