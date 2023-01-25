package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoziliang2019/blog-service/global"
	"github.com/zhaoziliang2019/blog-service/internal/service"
	"github.com/zhaoziliang2019/blog-service/pkg/app"
	"github.com/zhaoziliang2019/blog-service/pkg/errcode"
)

func GetAuth(ctx *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token": token})
}
