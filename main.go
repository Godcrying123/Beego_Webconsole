package main

import (
	_ "webconsole_sma/routers"

	"github.com/astaxie/beego"
)

// var (
// 	// the flag argument to customize the port num
// 	port = flag.Int("port", 8099, "the port for user to access")
// 	// set user should login the page or not
// 	LoginNeeded = flag.Bool("LoginNeeded", true, "set user should login the page or not")
// )

func main() {
	// flag.Parse()
	// beego.Run(":" + strconv.Itoa(*port))
	beego.Run()
}
