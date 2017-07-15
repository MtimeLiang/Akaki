package daos

import (
	"Akaki/beans"

	"github.com/astaxie/beego/orm"
)

func Insert(phone, password string) int {
	bean := beans.AdminBean{Phone: phone, Password: password}
	o := orm.NewOrm()
	// Insert返回的第一个参数返回主键id，不是SQL语句执行结果
	_, err := o.Insert(&bean)
	if err != nil {
		return 0
	}
	return 1
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
