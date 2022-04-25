package router

import (
	"user/config"

	"user/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initSwagger() {
	docs.SwaggerInfo.Title = "OJ SYS API"
	docs.SwaggerInfo.Description = "Gin API"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = config.ServerSetting.PrefixUrl
	docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func Router() *gin.Engine {
	gin.SetMode(config.ServerSetting.RunMode)
	r := gin.Default()

	// Swagger 配置
	go initSwagger()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 路由规则
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/users", func(ctx *gin.Context) {
			ctx.String(200, "get userinfo")
		})
		// userAPI := apiv1.Group("/users")
		// userAPI.Use()
		// {
		// 	problemAPI.GET("", service.GetProblemList)
		// 	problemAPI.GET("/:id", service.GetProblemDetail)
		// }
	}

	return r
}