package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goAPI/models"
)

type LoginContriller struct {
	beego.Controller
}

func (this *LoginContriller) Get() {
	fmt.Println("request:", this.Input().Get("name"))
	this.Data["json"] = "hello ray get "
	this.ServeJson()
}

func (this *LoginContriller) Post() {
	fmt.Println("request:", this.Input().Get("name"))
	name := this.Input().Get("name")
	password := this.Input().Get("password")

	account, err := models.QueryAccount(name, password)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = "login fail"
		return
	} else {
		fmt.Println("login success", account)
		this.Ctx.WriteString(account.Account)
		this.Data["json"] = "login success"
	}
	this.ServeJson()
}