package util

import (
	"github.com/L-Angel/LavBlog/models"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type InitConfig struct {
	config *models.SysConfig
}
const config_file_path="conf/sysconfig.js."
func (this *InitConfig)Load()(interface{},error){
	data,err:=ioutil.ReadFile(config_file_path)
	if(err!=nil){
		beego.Error(err)
		//beego.Info(ioutil.ReadDir("."))
		panic(err)
	}
	fmt.Println(data)
	v:=new(models.SysConfig)
	err=json.Unmarshal(data,v)
	if(err!=nil) {
		beego.Error(err)
		panic(err)
	}
	return v,nil
}

func (this *InitConfig)GetConfig()(interface{}){
	v,err:=this.Load()
	if(err!=nil){
		beego.Error(err)
		panic(nil)
	}
	fmt.Println(v)
	return v
}
