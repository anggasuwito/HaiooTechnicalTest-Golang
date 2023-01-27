package routes

import (
	"HaiooTechnicalTest-Golang/api/controller"
	"HaiooTechnicalTest-Golang/models"
	"github.com/gin-gonic/gin"
)

type CartRoutes struct {
	Route      *gin.Engine
	Controller controller.CartControllerI
}

type CartRoutesI interface {
	Routes()
}

func NewCartRoutes(route *gin.Engine, conf *models.Config) CartRoutesI {
	return &CartRoutes{
		Route:      route,
		Controller: controller.NewCartController(conf),
	}
}

func (r CartRoutes) Routes() {
	cartAPI := r.Route.Group("/cart")
	cartAPI.POST("", r.Controller.CreateProduct)
	cartAPI.DELETE("/:kodeProduk", r.Controller.DeleteProduct)
	cartAPI.GET("", r.Controller.GetProduct)
}
