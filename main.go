package main

import (
	"github.com/gin-gonic/gin"
	"deploy-station/app/routes"
)

func main() {
	r := gin.Default()

	routes.NewRoute(r)

	r.Run(":8010")
}
