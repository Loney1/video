package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["VideoApi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:BarrageController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           "/barrage/save",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:BarrageController"],
		beego.ControllerComments{
			Method:           "BarrageWs",
			Router:           "/barrage/ws",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:BaseController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:BaseController"],
		beego.ControllerComments{
			Method:           "ChannelRegion",
			Router:           "/channel/region",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:BaseController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:BaseController"],
		beego.ControllerComments{
			Method:           "ChannelType",
			Router:           "/channel/type",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:CommentController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           "/comment/list",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:CommentController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           "/comment/save",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:CommentController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:CommentController"],
		beego.ControllerComments{
			Method:           "SaveAll",
			Router:           "/comment/save/all",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:TopController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:TopController"],
		beego.ControllerComments{
			Method:           "ChannelTop",
			Router:           "/channel/top",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:TopController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:TopController"],
		beego.ControllerComments{
			Method:           "TypeTop",
			Router:           "/type/top",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "LoginDo",
			Router:           "/login/do",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SaveRegister",
			Router:           "/register/save",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:UserController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SendMessageDo",
			Router:           "/send/message",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelAdvert",
			Router:           "/channel/advert",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelHotList",
			Router:           "/channel/hot",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelRecommendRegionList",
			Router:           "/channel/recommend/region",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "GetChannelRecomendTypeList",
			Router:           "/channel/recommend/type",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelVideo",
			Router:           "/channel/video",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "UserVideo",
			Router:           "/user/video",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "VideoEpisodesList",
			Router:           "/video/episodes/list",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "VideoInfo",
			Router:           "/video/info",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "VideoSave",
			Router:           "/video/save",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "SaveAll",
			Router:           "/video/save/all",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "Search",
			Router:           "/video/search",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["VideoApi/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoApi/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "SendEs",
			Router:           "/video/send/es",
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
