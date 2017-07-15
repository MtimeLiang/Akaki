package routers

import (
	"Akaki/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/admin/signIn", &controllers.SignInController{})
}
