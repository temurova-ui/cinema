package main

import (
	"log"

	"github.com/temurova-ui/cinema/apigateway/api"
	// "github.com/temurova-ui/cinema/apigateway/config"
	// "github.com/temurova-ui/cinema/apigateway/services"
)

// @title Cinema Api
// @version v1.0.0
// @description Api for cinema project
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /

func main() {
	// conf, err := config.New("./config/config.env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// serviceManager, err := services.NewServiceManager(conf.Services)
	// if err != nil {
	// 	log.Fatalf("services.NewServiceManager(): %v", err)
	// }

	server := api.New(api.Option{
		// Conf:           *conf,
		// ServiceManager: serviceManager,
	})

	if err := server.Run(":8080" /*+ conf.HTTPPORT*/); err != nil {
		log.Fatalf("server.Run(): %v", err)
	}
}