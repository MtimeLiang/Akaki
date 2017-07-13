package services

import (
	"Akaki/beans"
	"Akaki/daos"
)

func Login(phone, password string) beans.AdminBean {
	bean := daos.Fetch(phone, password)
	return bean
}
