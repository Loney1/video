package routers

import (
	"video/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.VideoController{})
	beego.Include(&controllers.BaseController{})
	beego.Include(&controllers.CommentController{})
	beego.Include(&controllers.TopController{})
	beego.Include(&controllers.BarrageController{})
}
