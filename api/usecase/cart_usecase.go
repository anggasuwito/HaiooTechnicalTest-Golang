package usecase

import (
	"HaiooTechnicalTest-Golang/api/repository"
	"HaiooTechnicalTest-Golang/models"
)

type CartUsecase struct {
	Repo repository.CartRepositoryI
}

type CartUsecaseI interface {
}

func NewCartUsecase(conf *models.Config) CartUsecaseI {
	return &CartUsecase{
		Repo: repository.NewCartRepository(conf.DB),
	}
}
