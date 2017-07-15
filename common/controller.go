package common

/*
import (
	"Akaki/common/jwtbeego"
	"time"

	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

// @Title getToken
// @Description Get token from user session
// @Param	username		query 	string	true		"The username for get token"
// @Param	password		query 	string	true		"The password for get token"
// @Success 200 {string} Obtain Token
func (c *Controller) GetToken() {
	username := c.Ctx.Input.Query("username")
	password := c.Ctx.Input.Query("password")

	tokenString := ""
	if username == "admin" && password == "mipassword" {
		et := jwtbeego.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 3600, //Segundos
		}
		tokenString, _ = et.GetToken()
	}

	c.Data["json"] = "{'tokenString': '" + tokenString + "'}"
	c.ServeJSON()
	return
}

func (c *Controller) Prepare() {
	tokenString := c.Ctx.Input.Query("tokenString")

	et := jwtbeego.EasyToken{}
	valido, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}
*/
