package main

import (
	"github.com/gin-gonic/gin"
	"deploy-station/app/controllers"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./public/assets")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.LoadHTMLGlob("app/views/**/*")

	r.GET("/", controllers.Dashboard)
	r.GET("/home", controllers.Dashboard)
	r.GET("/item", controllers.ListItem)
	r.GET("/item/:name", controllers.ShowItem)
	r.GET("/item/:name/setting", controllers.SettingItem)
	r.POST("/item/:name/setting", controllers.SettingItem)
	r.GET("/push/:name", controllers.PreparePush)
	r.POST("/push/:name", controllers.ProcessPush)

	r.Run(":8090")
}
