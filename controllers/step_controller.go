package controllers

import (
	"github.com/astaxie/beego"
)

type StepController struct {
	beego.Controller
}

func (c *StepController) Get() {
	c.TplName = "step.html"
}
