package beans

import (
	"github.com/astaxie/beego/orm"
)

type AdminBean struct {
	Id       int    `orm:"auto"`
	Phone    string `orm:"size(20)"`
	Password string `orm:"size(20)"`
}

func init() {
	orm.RegisterModel(new(AdminBean))
}

/// 自定义表名
func (this *AdminBean) TableName() string {
	return "admin_member"
}
