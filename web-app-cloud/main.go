package main

import (
	controller "kubeedge-test/web-app-cloud/controller"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", new(controller.TrackController), "get:Index")
	beego.Router("/track/control/:trackId", new(controller.TrackController), "get,post:ControlTrack")

	beego.Run(":80")
}
