package main

import (
	_ "github.com/L-Angel/LavBlog/routers"
	"github.com/astaxie/beego"
	"github.com/L-Angel/LavBlog/controllers"
	"github.com/L-Angel/LavBlog/util"
)

func main() {
	sys:=new(util.InitSys)
	sys.Start()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

