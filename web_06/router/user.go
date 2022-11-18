package router

import (
	"github.com/gin-gonic/gin"
	"web_06/api"
)

func Initrouter() {
	r := gin.Default()
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.Run()
}
