package global

const (
	StatusOK            = 200
	StatusErr           = 400
	StatusLoginAuth     = 401
	StatusAuthErr       = 4001 // 未授权
	StatusMaxLimit      = 4002
	StatusSignErr       = 4003
	StatusReqParamErr   = 4005 // 请求参数错误
	StatusParamEmptyErr = 4006 // 请求参数位空
	StatusAuthIpErr     = 4007 // 白名单
	StatusOutTimeErr    = 4008 // 超时
	StatusDBErr         = 4009 // 数据库操作失败
	StatusNotActionErr  = 4010 // action 错误
)

var statusText = map[int]string{
	StatusOK:           "OK",
	StatusErr:          "error",
	StatusMaxLimit:     "You have reached maximum request limit.",
	StatusSignErr:      "sign error",
	StatusAuthErr:      "未授权",
	StatusAuthIpErr:    "IP 未授权",
	StatusOutTimeErr:   "request timeout",
	StatusNotActionErr: "Action Error",
	StatusLoginAuth:    "未登录",
}

func StatusAuthIpMsg(key string) string {
	return "IP: " + key + " 未授权"
}

func StatusAuthSourceMsg(key string) string {
	return "当前:" + key + "来源 未授权"
}

// StatusReqParamText 请求参数提示
func StatusReqParamText(key string) string {
	return "parameter request error : " + key
}

// StatusParamEmpty 参数不能为空
func StatusParamEmpty(key string) string {
	return key + "  为必填字段"
}

func StatusText(code int) string {
	return statusText[code]
}
