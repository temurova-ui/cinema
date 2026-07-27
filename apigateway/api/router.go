package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/temurova-ui/cinema/apigateway/docs"

	"github.com/temurova-ui/cinema/apigateway/api/handler"
	"github.com/temurova-ui/cinema/apigateway/config"
	"github.com/temurova-ui/cinema/apigateway/services"
)

type Option struct {
	Conf           config.Config
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	h := handler.NewHandler(option.ServiceManager)

	apiGroup := router.Group("/api")
	apiGroup.POST("/user/get", h.GetUser)
	apiGroup.POST("/movie/create", h.CreateMovie)
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}