package util

import (
	"Gin-Basic-Service/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// Response 返回结构体
type Response struct {
	Code    int         `json:"code"`    // 返回码
	Message string      `json:"message"` // 返回消息
	Data    interface{} `json:"data"`    // 返回数据
}

// PageResult 分页查询结果结构体
type PageResult struct {
	Total int64       `json:"total"` // 总记录数
	Items interface{} `json:"items"` // 查询结果列表
}

// SuccessResponse 成功返回
func SuccessResponse(data interface{}) *Response {
	return &Response{
		Code:    global.SuccessCode,
		Message: "Success",
		Data:    data,
	}
}

// ErrorResponse 错误返回
func ErrorResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// PaginationSuccess 分页查询成功返回
func PaginationSuccess(total int64, items interface{}) *Response {
	return &Response{
		Code:    global.SuccessCode,
		Message: "Success",
		Data: &PageResult{
			Total: total,
			Items: items,
		},
	}
}

// PageInfo 分页查询结构体
type PageInfo struct {
	Page       int         `json:"page" example:"1" `       // 当前页码
	PageSize   int         `json:"page_size" example:"10"`  // 每页记录数
	TotalCount int         `json:"total_count" example:"0"` // 总记录数
	QueryInfo  interface{} `json:"query_info"`              // 自定义查询条件
}

// GetOffset 根据当前页码和每页记录数计算偏移量
func (p *PageInfo) GetOffset() int {
	offset := (p.Page - 1) * p.PageSize
	if offset < 0 {
		return 0
	}
	return offset
}

// GetCountDB 返回获取count的数据查询
func (p *PageInfo) GetCountDB(db *gorm.DB) (count int64, err error) {
	err = db.Count(&count).Error
	return
}

// GetPageDB 返回获取分页的数据查询
func (p *PageInfo) GetPageDB(db *gorm.DB) {
	offset := p.GetOffset()
	limit := p.PageSize

	db.Limit(limit).Offset(offset)
}

// Respond 统一返回消息的方法
func Respond(c *gin.Context, code int, message string, data interface{}) {
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	// 根据返回消息结构体的返回码判断 HTTP 响应状态码
	switch code {
	case global.SuccessCode:
		c.JSON(http.StatusOK, response)
	case global.ErrorCode:
		c.JSON(http.StatusInternalServerError, response)
	case global.NotFoundCode:
		c.JSON(http.StatusNotFound, response)
	case global.Unauthorized:
		c.JSON(http.StatusUnauthorized, response)
	case global.InvalidRequest:
		c.JSON(http.StatusBadRequest, response)
	default:
		c.JSON(http.StatusOK, response)
	}
}
