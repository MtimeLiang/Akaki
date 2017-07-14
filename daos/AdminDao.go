package daos

import (
	"Akaki/beans"

	"github.com/astaxie/beego/orm"
)

func Insert(phone, password string) int {
	return 0
}

func Delete(phone, password string) int {
	return 0
}

func Update(phone, password string) int {
	return 0
}

func Fetch(phone, password string) beans.AdminBean {
	bean := beans.AdminBean{Phone: phone, Password: password}
	o := orm.NewOrm()
	o.Read(&bean, "phone", "password")
	return bean
}
