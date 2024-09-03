package usecase

import (
	"fmt"

	"github.com/rydhshlkhn/techtest-mirae/domain"
	"github.com/rydhshlkhn/techtest-mirae/repository"
)

type stockUsecase struct {
	stockRepo repository.StockRepository
}

type StockUsecase interface {
	CreateStock(stock *domain.Stock) (err error)
	GetAllStock() (stocks []domain.Stock)
	GetStockByID(id int) (stock *domain.Stock, err error)
	UpdateStock(stock *domain.Stock) (err error)
	DeleteStock(id int) (err error)
}

func NewStockUsecase(repo repository.StockRepository) StockUsecase {
	return &stockUsecase{stockRepo: repo}
}

func (u *stockUsecase) CreateStock(stock *domain.Stock) (err error) {
	if err := u.stockRepo.CreateStock(stock); err != nil {
		return fmt.Errorf("failed to  create stock: %v", err)
	}

	return
}

func (u *stockUsecase) GetAllStock() (stocks []domain.Stock) {
	return u.stockRepo.GetAllStock()
}

func (u *stockUsecase) GetStockByID(id int) (stock *domain.Stock, err error) {
	return u.stockRepo.GetStockByID(id)
}

func (u *stockUsecase) UpdateStock(stock *domain.Stock) (err error) {
	return u.stockRepo.UpdateStock(stock)
}

func (u *stockUsecase) DeleteStock(id int) (err error) {
	stock, err := u.stockRepo.GetStockByID(id)
	if err != nil {
		return
	}
	return u.stockRepo.DeleteStock(stock.ID)
}
