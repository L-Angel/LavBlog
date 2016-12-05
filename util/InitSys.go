package util

import "github.com/astaxie/beego"

type InitSys struct {
	config *InitConfig

}
func (this *InitSys)init(){
	beego.SetLogger("file",`{"filename":"logs/logger.log"}`)
	beego.SetLogFuncCall(true)
	beego.Informational("App Start")
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
		}
	}()
	return
}
func (this *InitSys)Start(){
	this.init()
	this.config=new(InitConfig)
	this.config.GetConfig()
	return
}