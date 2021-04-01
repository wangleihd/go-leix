package controllers

import (
	"test/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type NewsController struct {
	beego.Controller
}

type ClassController struct {
	beego.Controller
}

type Res struct {
	code    int
	message string
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *NewsController) GetAll() {
	page, _ := u.GetInt("page")

	if page <= 0 {
		page = 1
	}

	info := models.GetAllNewsEn(page)
	msg := map[string]interface{}{"code": 0, "msg": "success", "data": info}
	u.Data["json"] = msg
	u.ServeJSON()
}

func (u *ClassController) GetAll() {
	name := u.GetString("name")
	page, _ := u.GetInt("page")

	if name == "" {
		name = "us"
	}
	if page <= 0 {
		page = 1
	}
	info := models.GetClass(name, page)
	msg := map[string]interface{}{"code": 0, "msg": "success", "data": info}
	u.Data["json"] = msg
	u.ServeJSON()
}
