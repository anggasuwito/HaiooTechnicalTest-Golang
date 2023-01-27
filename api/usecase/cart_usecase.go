package usecase

import (
	"HaiooTechnicalTest-Golang/api/repository"
	"HaiooTechnicalTest-Golang/models"
	"errors"
	"github.com/google/uuid"
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
	cart, err := u.Repo.FindByName(body.NamaProduk)
	if err != nil {
		return err
	}

	if cart.KodeProduk != "" {
		cart.Kuantitas += body.Kuantitas
		return u.Repo.UpdateProduct(cart)
	} else {
		return u.Repo.CreateProduct(models.Cart{
			KodeProduk: uuid.New().String(),
			NamaProduk: body.NamaProduk,
			Kuantitas:  body.Kuantitas,
		})
	}
}

func (u CartUsecase) DeleteProduct(kodeProduk string) error {
	_, err := uuid.Parse(kodeProduk)
	if err != nil {
		return errors.New("invalid uuid format")
	}
	data, err := u.Repo.FindByCode(kodeProduk)
	if err != nil {
		return err
	}
	if data.KodeProduk == "" {
		return errors.New("record not found")
	}
	err = u.Repo.DeleteProduct(kodeProduk)
	return err
}

func (u CartUsecase) GetProduct(params models.CartRequestParameter) ([]models.Cart, error) {
	res, err := u.Repo.GetProduct(params)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("record not found")
	}
	return res, nil
}
