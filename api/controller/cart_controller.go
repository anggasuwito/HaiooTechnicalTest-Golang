package controller

import (
	"HaiooTechnicalTest-Golang/api/usecase"
	"HaiooTechnicalTest-Golang/models"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	Usecase usecase.CartUsecaseI
}

type CartControllerI interface {
	CreateProduct(g *gin.Context)
	DeleteProduct(g *gin.Context)
	GetProduct(g *gin.Context)
}

func NewCartController(conf *models.Config) CartControllerI {
	return &CartController{
		Usecase: usecase.NewCartUsecase(conf),
	}
}

func (c CartController) CreateProduct(g *gin.Context) {

}

func (c CartController) DeleteProduct(g *gin.Context) {

}

func (c CartController) GetProduct(g *gin.Context) {

}
