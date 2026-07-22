package handler

import(
	"github.com/temurova-ui/cinema/apigateway/services"
)

type handler struct{
	serviceManager services.IServiceManager
}

func NewHandler (serviceManager services.IServiceManager)*handler{
	return &handler{
		serviceManager: serviceManager,
	}
}