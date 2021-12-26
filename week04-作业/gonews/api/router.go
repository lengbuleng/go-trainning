package api

import (
	v1 "app/api/v1"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {

	api := g.Group("/api/v1/")
	middleware.Cors(api)
	api.Use(middleware.CustomError)
	middleware.RegisterPPROF(g, "/dev/pprof")

	api.POST("login", v1.Login)
	api.POST("register", v1.Register)

	api.Use(gin.Logger())

	api.Use(middleware.Auth())

	api.GET("authTest", v1.AuthTest)
	api.POST("addNews", v1.AddNews)
	api.GET("section", v1.GetSection)

}
