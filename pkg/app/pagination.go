package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoziliang2019/blog-service/global"
	"github.com/zhaoziliang2019/blog-service/pkg/convert"
)

func GetPage(ctx *gin.Context) int {
	page := convert.StrTo(ctx.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}
func GetPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.Appsetting.DefaultPageSize
	}
	if pageSize > global.Appsetting.MaxPageSize {
		return global.Appsetting.MaxPageSize
	}
	return pageSize
}
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
