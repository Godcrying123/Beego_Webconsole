package controllers

type MachineController struct {
	BaseController
}

func (this *MachineController) Get() {
	this.TplName = "machine.html"
}
