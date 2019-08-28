package controllers

import (
	"github.com/astaxie/beego"
)

type STFPController struct {
	beego.Controller
}

func (this *STFPController) Get() {
	this.TplName = "file.html"

}
