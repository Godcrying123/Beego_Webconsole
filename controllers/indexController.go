package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func init() {

}

func (c *IndexController) Get() {
	c.TplName = "index.html"
}
