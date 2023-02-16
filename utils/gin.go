package utils

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// GinContextToContext gin.context 转换为 context.Contex
func GinContextToContext(ctx *gin.Context) context.Context {
	return ctx.Request.Context()
}

// GinGetRawData gin 获取json格式数据
func GinGetRawData(ctx *gin.Context, data interface{}) (err error) {
	body, err := ctx.GetRawData()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)

	return err
}

// GinBindJSON gin获取json格式数据
func GinBindJSON(ctx *gin.Context, data interface{}) (err error) {
	err = ctx.BindJSON(&data)

	return err
}

// GinPostForm gin获取post表单提交数据
func GinPostForm(ctx *gin.Context, key string) (out string) {

	if ctx.PostForm(key) != "" {
		return ctx.PostForm(key)
	}

	return ctx.Request.URL.Query().Get(key)
}

func GinUrlPath(ctx *gin.Context) (out string) {

	return ctx.Request.URL.Path
}

func GetGinContextValue(ctx *gin.Context, key interface{}) interface{} {
	return ctx.Request.Context().Value(key)
}

func SetGinContextValue(ctx *gin.Context, key interface{}, value interface{}) {
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), key, value))
}
