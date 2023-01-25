package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoziliang2019/blog-service/global"
	"github.com/zhaoziliang2019/blog-service/internal/model"
	"github.com/zhaoziliang2019/blog-service/internal/service"
	"github.com/zhaoziliang2019/blog-service/pkg/app"
	"github.com/zhaoziliang2019/blog-service/pkg/convert"
	"github.com/zhaoziliang2019/blog-service/pkg/errcode"
)

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

func NewArticle() Article {
	return Article{}
}
func (t Article) Get(ctx *gin.Context) {
	param := service.ArticleRequest{
		ID: convert.StrTo(ctx.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(ctx)
	valid, err := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	article, errs := svc.GetArticle(&param)
	if errs != nil {
		global.Logger.Errorf("svc.GetArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail.WithDetails(err.Errors()...))
		return
	}
	response.ToResponse(article)
	return
}
func (t Article) List(ctx *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(ctx),
		PageSize: app.GetPageSize(ctx),
	}
	articles, totalRows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	response.ToResponseList(articles, totalRows)
	return
}
func (t Article) Create(ctx *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle errs:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (t Article) Update(ctx *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(ctx.Param("id")).MustUInt32()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle errs:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (t Article) Delete(ctx *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(ctx.Param("id")).MustUInt32()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
