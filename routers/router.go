package routers

import (
	"github.com/L-Angel/LavBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
