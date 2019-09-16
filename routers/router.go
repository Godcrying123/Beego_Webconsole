package routers

import (
	"webconsole_sma/controllers"
	_ "webconsole_sma/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/machine", &controllers.MachineController{})
	beego.Router("/machine/import", &controllers.MachineController{}, "post:PostMachines")
	beego.Router("/node/:hostname([\\w]+)", &controllers.IndexController{})
	beego.Router("/step", &controllers.StepController{})
	// beego.Router("host/", &controllers.HostController{})
	beego.Router("host/ws", &controllers.HostWebSocketController{})
	// beego.Router("host/", &controllers.HostController{}, "post:StopHostSync")
	beego.Router("ssh/ws", &controllers.SSHWebSocketController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/service/ws", &controllers.ServiceWebSocketController{})
	beego.Router("/file/*", &controllers.FileController{})
	// beego.Router("/sftp/file/", &controllers.STFPController{}, "post:SaveAndFind")
	// beego.Router("/node/:hostname([\\w]+)/file/", &controllers.FileController{}, "get:GetSFTP")
	beego.Router("/node/:hostname/file/*", &controllers.STFPController{})
	beego.Router("/node/:hostname/file/", &controllers.STFPController{})
	beego.Router("/task", &controllers.TaskController{})
}
