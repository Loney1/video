package routers

import (
	"Video/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/channel", &controllers.MainController{}, "get:Channel")
	beego.Router("/channel/video/data", &controllers.MainController{}, "*:ChannelVideoData")
	beego.Router("/top", &controllers.MainController{}, "get:Top")
	beego.Router("/show", &controllers.MainController{}, "get:Show")

	beego.Router("/comment/list", &controllers.MainController{}, "post:GetCommentList")
	beego.Router("/comment/save", &controllers.MainController{}, "post:SaveComment")

	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/mini/login", &controllers.UserController{}, "get:MiniLogin")
	beego.Router("/login/do", &controllers.UserController{}, "post:LoginDo")
	beego.Router("/register", &controllers.UserController{}, "get:Register")
	beego.Router("/register/save", &controllers.UserController{}, "post:RegisterSave")
	beego.Router("/ucenter/video", &controllers.UserController{}, "get:UserVideo")
	beego.Router("/my/video", &controllers.UserController{}, "post:GetMyVideos")

	beego.Router("/send/message", &controllers.UserController{}, "get:SendMessage")
	beego.Router("/send/message/do", &controllers.UserController{}, "post:SendMessageDo")

	beego.Router("/barrage/save", &controllers.MainController{}, "post:SaveBarrage")

	beego.Router("/upload", &controllers.UserController{}, "get:Upload")
	beego.Router("/uploadDo", &controllers.UserController{}, "post:UploadVideo")
	beego.Router("/uploadInfoDo", &controllers.UserController{}, "post:UploadInfoDo")

	beego.Router("/search", &controllers.MainController{}, "get:Search")
	beego.Router("/search/data", &controllers.MainController{}, "*:SearchData")
}
