package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoziliang2019/blog-service/global"
	"github.com/zhaoziliang2019/blog-service/internal/service"
	"github.com/zhaoziliang2019/blog-service/pkg/app"
	"github.com/zhaoziliang2019/blog-service/pkg/convert"
	"github.com/zhaoziliang2019/blog-service/pkg/errcode"
	"github.com/zhaoziliang2019/blog-service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}
func (u Upload) UploadFile(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	file, fileHeader, err := ctx.Request.FormFile("file")
	fileType := convert.StrTo(ctx.PostForm("type")).MustUInt32()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err:%v", err)
		response.ToErrorResponse(errcode.ERROR_UPLOAD_FILE_FAIL.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
