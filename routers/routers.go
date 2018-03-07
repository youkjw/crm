package routers

import (
	//"github.com/astaxie/beego/context"
	"github.com/crm/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.IndexController{})

	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.MainController{})



	//beego.Router("/", &controllers.MainController{})

	//beego.Get("/", func(ctx *context.Context){
	//	ctx.Output.Body([]byte("hello world"))
	//})

	//beego.Router("/api/?:id", &controllers.RController{})

}