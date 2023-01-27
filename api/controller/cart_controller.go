package controller

import (
	"HaiooTechnicalTest-Golang/api/usecase"
	"HaiooTechnicalTest-Golang/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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
	var body models.CartRequestBody
	err := g.ShouldBindJSON(&body)
	if err != nil {
		var v validator.ValidationErrors
		if errors.As(err, &v) {
			for _, fieldError := range v {
				g.JSON(http.StatusBadRequest, models.Response{
					Message: fmt.Sprintf("%v %+v", fieldError.Field(), fieldError.Tag()),
				})
				return
			}
		}
		g.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
		return
	}
	err = c.Usecase.CreateProduct(body)
	if err != nil {
		g.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, models.Response{
		Message: "Success",
	})
}

func (c CartController) DeleteProduct(g *gin.Context) {
	kodeProduk := g.Param("kodeProduk")
	err := c.Usecase.DeleteProduct(kodeProduk)
	if err != nil {
		g.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, models.Response{
		Message: "Success",
	})
}

func (c CartController) GetProduct(g *gin.Context) {
	var params models.CartRequestParameter
	err := g.ShouldBindQuery(&params)
	if err != nil {
		g.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
		return
	}
	res, err := c.Usecase.GetProduct(params)
	if err != nil {
		g.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    res,
	})
}
