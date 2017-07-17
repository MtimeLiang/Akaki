package common

import (
	"Akaki/common/macro"
	"Akaki/common/token"
	"time"

	"github.com/astaxie/beego/context"
)

/// Verify token
func VerifyToken(ctx *context.Context) bool {
	ts := ctx.GetCookie(macro.CookieName)
	if len(ts) == 0 {
		return false
	}
	t, err := token.Parse(ts, token.SignKey)
	if t.Valid && err == nil {
		c, ok := token.GetMapClaims(t)
		exp := c["exp"]
		now := time.Now()
		if ok {
			if exp.(time.Time).Before(now) == true {
				return true
			}
		}
	}
	return false
}
