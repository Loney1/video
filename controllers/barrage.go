package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"net/http"
)

type BarrageController struct {
	beego.Controller
}

type WsData struct {
	CurrentTime int
	EpisodesId  int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//获取弹幕websocket
// @router /barrage/ws [*]
func (this *BarrageController) BarrageWs() {

}
