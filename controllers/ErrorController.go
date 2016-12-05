package controllers

type ErrorController struct {
	baseController
}
func (this *ErrorController)Error404(){
	this.TplName="Error404.tpl"
}
func (this *ErrorController)Error500(){
	this.TplName="Error500.tpl"

}
func (this *ErrorController)ErrorDb(){
	this.TplName="ErrorDb.tpl"
}
