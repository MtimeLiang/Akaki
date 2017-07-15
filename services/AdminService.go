package services

import (
	"Akaki/beans"
	"Akaki/daos"
)

func SignIn(phone, password string) int {
	result := daos.Insert(phone, password)
	return result
}

func Login(phone, password string) beans.AdminBean {
	bean := daos.Fetch(phone, password)
	return bean
}
