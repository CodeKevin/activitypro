package main

import (
	. "activitypro/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.POST("/activity", AddActivityApi)
	router.POST("/user", AddUserApiTest)
	router.GET("/activity", GetAllPersonApi)
	return router
}
