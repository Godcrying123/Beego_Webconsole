package controllers

import (
	"github.com/astaxie/beego"
)

//The Init Definition for main controller
type MainController struct {
	beego.Controller
}

//The Init Definition for main controller's Get method
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
