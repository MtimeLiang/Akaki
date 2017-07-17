package common

import (
	"github.com/astaxie/beego/context"
)

func ConfigResponseHeader(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
}
