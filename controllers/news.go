package controllers

import (
	"test/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type NewsController struct {
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
	info := models.GetAllNewsEn()
	msg := map[string]interface{}{"code": 0, "msg": "success", "data": info}
	u.Data["json"] = msg
	u.ServeJSON()
}
