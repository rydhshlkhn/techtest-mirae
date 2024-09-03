package repository

import (
	"github.com/rydhshlkhn/techtest-mirae/domain"
	"gorm.io/gorm"
)

type stockRepository struct {
	db *gorm.DB
}

type StockRepository interface {
	CreateStock(stock *domain.Stock) (err error)
	GetStockByID(id int) (stock *domain.Stock, err error)
	GetAllStock() (stocks []domain.Stock)
	UpdateStock(stock *domain.Stock) (err error)
	DeleteStock(id int) (err error)
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) CreateStock(stock *domain.Stock) (err error) {
	return r.db.Create(stock).Error
}

func (r *stockRepository) GetAllStock() (stocks []domain.Stock) {
	r.db.Find(&stocks)
	return
}

func (r *stockRepository) GetStockByID(id int) (stock *domain.Stock, err error) {
	err = r.db.Where("id = ?", id).First(&stock).Error
	return
}

func (r *stockRepository) UpdateStock(stock *domain.Stock) (err error) {
	err = r.db.Where("id = ?", stock.ID).Updates(&stock).Error
	return
}

func (r *stockRepository) DeleteStock(id int) (err error) {
	var stock domain.Stock
	err = r.db.Where("id = ?", id).Delete(stock).Error
	return
}
