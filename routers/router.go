package routers

import (
	"webconsole_sma/controllers"
	_ "webconsole_sma/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/host", &controllers.HostController{})
	beego.AutoRouter(&controllers.StepController{})
	beego.Router("/step", &controllers.StepController{})
	beego.AutoRouter(&controllers.ServiceController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/file/*", &controllers.FileController{})
	// beego.AutoRouter(&controllers.FileController{})
}
