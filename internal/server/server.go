package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/koustubh25/babybucks/internal/handlers"
)

func RunServer() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/fetchbalancesandgoals", handlers.FetchBalancesAndGoals)
	r.NoRoute(static.Serve("/", static.LocalFile("front/dist", false)))

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
