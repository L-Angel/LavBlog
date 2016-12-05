package v1

import "github.com/astaxie/beego"

func init()  {
	beego.NSNamespace("/v1",
	beego.Include("/login"))
}
