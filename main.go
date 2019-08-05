package main

import (
	_ "webconsole_sma/routers"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

func main() {
	//beego.Run()
	strout, err := utils.Command_Exec("df -ah")
	beego.Info(strout)
	beego.Error(err)
}
