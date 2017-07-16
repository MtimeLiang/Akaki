package controllers

import (
	"Akaki/common"
	"Akaki/services"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type SignInController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	phone := c.GetString("phone")
	password := c.GetString("password")
	// MD5加密
	password = common.MD5(password)

	r := make(map[string]interface{})
	defer func() {
		c.Data["json"] = r
		c.ServeJSON()
	}()

	if len(phone) > 0 && len(password) > 0 {
		bean := services.Login(phone, password)
		if bean.Id > 0 {
			r["infoMsg"] = "登录成功"
			r["status"] = 1
			r["data"] = bean
			return
		}
		r["infoMsg"] = "登录失败，账号或密码输入错误"
		r["status"] = 0
		r["data"] = new(map[string]interface{})
		return
	}
	r["infoMsg"] = "登录失败，手机号和密码不能为空"
	r["status"] = 0
	r["data"] = new(map[string]interface{})
	return
}

func (c *SignInController) Post() {
	phone := c.GetString("phone")
	password := c.GetString("password")
	// MD5加密
	password = common.MD5(password)

	r := make(map[string]interface{})
	defer func() {
		c.Data["json"] = r
		c.ServeJSON()
	}()

	if len(phone) > 0 && len(password) > 0 {
		bean := services.Login(phone, password)
		if bean.Id > 0 {
			r["infoMsg"] = "注册失败，该用户已存在"
			r["status"] = 0
			r["data"] = make(map[string]interface{})
			return
		}
		result := services.SignIn(phone, password)
		if result == 1 {
			r["infoMsg"] = "注册成功"
			r["status"] = 1
			r["data"] = make(map[string]interface{})
			return
		}
	}
	r["infoMsg"] = "注册失败，手机号和密码不能为空"
	r["status"] = 0
	r["data"] = new(map[string]interface{})
	return
}
