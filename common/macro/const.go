package macro

const (
	CookieName = "_Akaki_"
)

type ResInfo struct {
	Status  int         `json:"status"`
	InfoMsg string      `json:"infoMsg"`
	Data    interface{} `json:"data"`
}

var (
	ErrOperate       = &ResInfo{Status: 2, InfoMsg: "操作失败", Data: ""}
	ErrTokenInvalid  = &ResInfo{Status: 3, InfoMsg: "登录过期，请重新登录", Data: ""}
	ErrArgsNotCorret = &ResInfo{Status: 4, InfoMsg: "参数错误", Data: ""}
)
