package controllers

import (
	"Akaki/common"
	"Akaki/services"
	"time"

	"Akaki/common/macro"
	"Akaki/common/token"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
)

type LoginController struct {
	beego.Controller
}

type SignInController struct {
	beego.Controller
}

type LogoutController struct {
	beego.Controller
}

type GetAdminInfoController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	common.ConfigResponseHeader(c.Ctx)
	phone := c.GetString("phone")
	password := c.GetString("password")
	// MD5加密
	password = common.MD5(password)

	if len(phone) == 0 || len(password) == 0 {
		c.Data["json"] = macro.ResInfo{InfoMsg: "登录失败，手机号和密码不能为空", Status: 0, Data: ""}
		c.ServeJSON()
	}
	bean := services.Login(phone, password)
	if bean.Id == 0 {
		c.Data["json"] = macro.ResInfo{InfoMsg: "登录失败，账号或密码输入错误", Status: 0, Data: ""}
		c.ServeJSON()
	}
	c.Data["json"] = macro.ResInfo{InfoMsg: "登录成功", Status: 1, Data: bean}
	// 注入Cookie
	cal := jwt.MapClaims{
		"phone":    bean.Phone,
		"password": bean.Password,
		"userId":   bean.Id,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	t, _ := token.New(token.SignKey, cal)
	c.Ctx.SetCookie(macro.CookieName, t)
	c.ServeJSON()
}

func (c *SignInController) Post() {
	phone := c.GetString("phone")
	password := c.GetString("password")
	// MD5加密
	password = common.MD5(password)

	if len(phone) == 0 || len(password) == 0 {
		c.Data["json"] = macro.ResInfo{InfoMsg: "注册失败，手机和密码不能为空", Status: 0, Data: ""}
		c.ServeJSON()
	}
	bean := services.Login(phone, password)
	if bean.Id == 0 {
		result := services.SignIn(phone, password)
		if result == 1 {
			c.Data["json"] = macro.ResInfo{InfoMsg: "注册成功", Status: 1, Data: ""}
			// 注入Cookie
			cal := jwt.MapClaims{
				"phone":    bean.Phone,
				"password": bean.Password,
				"userId":   bean.Id,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			}
			t, _ := token.New(token.SignKey, cal)
			c.Ctx.SetCookie(macro.CookieName, t)
			c.ServeJSON()
		}
	}
	c.Data["json"] = macro.ResInfo{InfoMsg: "注册失败", Status: 0, Data: ""}
	c.ServeJSON()
}

func (c *LogoutController) Get() {
	ts := c.Ctx.GetCookie(macro.CookieName)
	t, err := token.Parse(ts, token.SignKey)
	if err != nil || !t.Valid {
		c.Data["json"] = macro.ResInfo{InfoMsg: "token验证失败", Status: 0, Data: ""}
		c.ServeJSON()
	}
	c.Ctx.SetCookie(macro.CookieName, "")
	c.Data["json"] = macro.ResInfo{InfoMsg: "退出登录成功", Status: 1, Data: ""}
	c.ServeJSON()
}

func (c *GetAdminInfoController) Get() {
	ts := c.Ctx.GetCookie(macro.CookieName)
	tr, _ := token.Parse(ts, token.SignKey)
	if tr != nil {
		m, _ := token.GetMapClaims(tr)
		phone := m["phone"]
		password := m["password"]
		bean := services.Login(phone.(string), password.(string))
		if bean.Id > 0 {
			c.Data["json"] = macro.ResInfo{InfoMsg: "获取用户信息成功", Status: 1, Data: bean}
			c.ServeJSON()
		}
	}
	c.Data["json"] = macro.ResInfo{InfoMsg: "获取用户信息失败", Status: 0, Data: ""}
	c.ServeJSON()
}
