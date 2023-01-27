package usecase

import (
	"HaiooTechnicalTest-Golang/api/repository"
	"HaiooTechnicalTest-Golang/models"
	"github.com/google/uuid"
	"time"
)

type CartUsecase struct {
	Repo repository.CartRepositoryI
}

type CartUsecaseI interface {
	CreateProduct(body models.CartRequestBody) error
	DeleteProduct(kodeProduk string) error
	GetProduct(params models.CartRequestParameter) ([]models.Cart, error)
}

func NewCartUsecase(conf *models.Config) CartUsecaseI {
	return &CartUsecase{
		Repo: repository.NewCartRepository(conf.DB),
	}
}

func (u CartUsecase) CreateProduct(body models.CartRequestBody) error {
	return u.Repo.CreateProduct(models.Cart{
		KodeProduk: uuid.New().String(),
		NamaProduk: body.NamaProduk,
		Kuantitas:  body.Kuantitas,
		CreatedAt:  time.Now(),
	})
}

func (u CartUsecase) DeleteProduct(kodeProduk string) error {
	return u.Repo.DeleteProduct(kodeProduk)
}

func (u CartUsecase) GetProduct(params models.CartRequestParameter) ([]models.Cart, error) {
	return u.Repo.GetProduct(params)
}
