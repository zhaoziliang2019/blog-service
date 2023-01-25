package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(1000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1000003, "鉴权失败,找不到对应的Appkey和Appsecret")
	UnauthorizedTokenError    = NewError(1000004, "鉴权失败,Token错误")
	UnauthorizedTokenTimeout  = NewError(1000005, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(1000006, "鉴权失败,Token生成失败")
	TooManyRequests           = NewError(1000007, "请求过多")
	ErrorGetTagListFail       = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail        = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail        = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail        = NewError(20010004, "删除标签失败")
	ErrorCountTagFail         = NewError(20010005, "统计标签失败")
	ErrorGetArticleFail       = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail      = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail    = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail    = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail    = NewError(20020005, "删除文章失败")
	ErrorCountArticleFail     = NewError(20020006, "统计文章失败")
	ERROR_UPLOAD_FILE_FAIL    = NewError(20030001, "上传文件失败")
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d 已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &Error{
		code:    code,
		msg:     msg,
		details: nil,
	}
}
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d,错误信息：%s", e.Code(), e.Msg())
}
func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}
func (e *Error) Details() []string {
	return e.details
}
func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}
func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.code:
		fallthrough
	case UnauthorizedTokenError.code:
		fallthrough
	case UnauthorizedTokenGenerate.code:
		fallthrough
	case UnauthorizedTokenTimeout.code:
		return http.StatusUnauthorized
	case TooManyRequests.code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
