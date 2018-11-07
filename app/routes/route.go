package routes

import (
	"github.com/gin-gonic/gin"
	"deploy-station/app/controllers"
	"deploy-station/app/middleware"
)

func NewRoute(r *gin.Engine) {
	r.Static("/assets", "./public/assets")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.LoadHTMLGlob("app/views/**/*")

	r.GET("/login", controllers.Login)
	r.POST("/login", controllers.Login)
	r.POST("/item/:name/setting", controllers.SettingItem)
	r.POST("/push/:name", controllers.ProcessPush)

	httpGroup := r.Group("/")
	httpGroup.Use(middleware.UserAuth())
	{
		httpGroup.GET("/", controllers.Dashboard)
		httpGroup.GET("/home", controllers.Dashboard)
		httpGroup.GET("/item", controllers.ListItem)
		httpGroup.GET("/item/:name", controllers.ShowItem)
		httpGroup.GET("/item/:name/setting", controllers.SettingItem)
		httpGroup.GET("/push/:name", controllers.PreparePush)
	}
}
