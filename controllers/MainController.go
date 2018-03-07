package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	param := this.GetString
	fmt.Println(param)
	//this.Ctx.WriteString("Main")
}

func (this *MainController) Add() {
	this.Ctx.WriteString("Main/Add")
}