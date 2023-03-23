package global

import (
	"encoding/json"
	"net/http"

	"github.com/ct-llper/go-pkg/plugin/metadata"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

// Body 内容
type Body struct {
	Action string                 `json:"action"`
	Sign   string                 `json:"sign"`
	Data   map[string]interface{} `json:"data"`
	Source []int                  `json:"source"` // 来源
	Time   int64                  `json:"time"`
}

// Response 返回状态
type Response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// ResponseData 返回数据
type ResponseData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseOkData 成功-带数据
func ResponseOkData(c *gin.Context, data interface{}) {
	res := ResponseData{
		Code: StatusOK,
		Msg:  StatusText(StatusOK),
		Data: data,
	}
	response, _ := json.Marshal(res)
	logx.WithContext(metadata.GetContext(c)).Info("Response:" + string(response))
	c.JSON(http.StatusOK, res)
}

// ResponseOk 成功-不带数据
func ResponseOk(c *gin.Context) {
	res := Response{
		Code: StatusOK,
		Msg:  StatusText(StatusOK),
	}
	response, _ := json.Marshal(res)
	logx.WithContext(metadata.GetContext(c)).Info("Response:" + string(response))
	c.JSON(http.StatusOK, res)
}

// ResponseError 失败返回、失败描述
func ResponseError(c *gin.Context, err error) {
	res := Response{
		Code: StatusErr,
		Msg:  err.Error(),
	}
	response, _ := json.Marshal(res)
	logx.WithContext(metadata.GetContext(c)).Info("Response:" + string(response))
	c.JSON(http.StatusOK, res)
}

// ResponseStatus 自定义code  msg
func ResponseStatus(c *gin.Context, res *Response) {
	response, _ := json.Marshal(res)
	logx.WithContext(metadata.GetContext(c)).Info("Response:" + string(response))
	c.JSON(http.StatusOK, res)
}

type WechatResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ResponseWechat(c *gin.Context, res *WechatResponse) {
	response, _ := json.Marshal(res)
	logx.WithContext(metadata.GetContext(c)).Info("ResponseWechat:" + string(response))
	c.JSON(http.StatusOK, res)
}

// ApiError api错误
type ApiError struct {
	Name      string `json:"name"`       // 错误名称
	ErrorCode string `json:"error_code"` // 错误状态码
	Content   string `json:"content"`    // 错误内容
	Param     string `json:"param"`      // 请求参数
	CreateAt  int64  `json:"create_at"`  // 创建时间
}
