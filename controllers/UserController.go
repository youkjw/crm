package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Ctx.WriteString("User")
}

func (this *UserController) Add() {
	this.Ctx.WriteString("User/Add")
}