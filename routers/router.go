package routers

import (
	"github.com/astaxie/beego"
	"terranova/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.ClusterController{})
	beego.Include(&controllers.UserController{})
}
