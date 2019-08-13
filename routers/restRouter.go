package routers

import (
	"webconsole_sma/controllers"
	_ "webconsole_sma/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/index", &controllers.IndexController{})
	beego.Router("/api/service", &controllers.ServiceController{})
	beego.Router("/api/host", &controllers.HostController{})
	beego.Router("/api/step", &controllers.StepController{})
	beego.Router("/api/file", &controllers.FileController{})
}
