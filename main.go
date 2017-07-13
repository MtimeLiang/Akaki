package main

import (
	_ "Akaki/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 连接数据库的DSN格式如下
	// username:password@protocol(address)/dbname?param=value
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/perfect_note?charset=utf8")
}

func main() {
	beego.Run()
}
