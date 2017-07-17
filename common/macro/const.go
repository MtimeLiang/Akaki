package macro

type ResError struct {
	Status  int         `json:"status"`
	InfoMsg string      `json:"infoMsg"`
	Data    interface{} `json:"data"`
}

var (
	ErrOperate       = &ResError{Status: 2, InfoMsg: "操作失败", Data: ""}
	ErrTokenInvalid  = &ResError{Status: 3, InfoMsg: "登录过期，请重新登录", Data: ""}
	ErrArgsNotCorret = &ResError{Status: 4, InfoMsg: "参数错误", Data: ""}
)
