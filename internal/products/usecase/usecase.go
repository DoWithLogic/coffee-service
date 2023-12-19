package usecase

import "github.com/DoWithLogic/coffee-service/internal/products"

type usecase struct{}

func NewUseCase() products.Usecase {
	return &usecase{}
}
