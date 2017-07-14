package controllers

import (
	"Akaki/services"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type SignInController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	phone := c.GetString("phone")
	password := c.GetString("password")

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

}
