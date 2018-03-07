package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type Person struct {
	UserName string
}

func (this *IndexController) Index() {

	var nov = [...]int{1, 2, 3}
	this.Data["nov"] = nov
	beego.Informational(nov)
	this.TplName = "index.html"
}

func (this *IndexController) Get() {
	fmt.Println(this.Data)
	this.Ctx.WriteString("Index")
}

func (this *IndexController) Add() {

}