package routers

import (
	"webconsole_sma/controllers"
	_ "webconsole_sma/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/host", &controllers.HostController{})
	beego.Router("/step", &controllers.StepController{})
}
