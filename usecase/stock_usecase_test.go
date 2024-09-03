package usecase

import (
	"errors"
	"testing"

	"github.com/rydhshlkhn/techtest-mirae/domain"
	mocks "github.com/rydhshlkhn/techtest-mirae/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	errFoo = errors.New("something error")
)

func Test_CreateStock(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		stockRepo := &mocks.StockRepository{}
		stockRepo.On("CreateStock", mock.Anything).Return(nil)

		uc := NewStockUsecase(stockRepo)

		err := uc.CreateStock(&domain.Stock{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		stockRepo := &mocks.StockRepository{}
		stockRepo.On("CreateStock", mock.Anything).Return(errFoo)

		uc := NewStockUsecase(stockRepo)

		err := uc.CreateStock(&domain.Stock{})
		assert.Error(t, err)
	})
}

func Test_GetAllStock(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		stockRepo := &mocks.StockRepository{}
		stockRepo.On("GetAllStock", mock.Anything).Return([]domain.Stock{})

		uc := NewStockUsecase(stockRepo)

		res := uc.GetAllStock()
		assert.NotNil(t, res)
	})
}

func Test_GetStockByID(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		stockRepo := &mocks.StockRepository{}
		stockRepo.On("GetStockByID", mock.Anything).Return(&domain.Stock{}, nil)

		uc := NewStockUsecase(stockRepo)

		_, err := uc.GetStockByID(1)
		assert.NoError(t, err)
	})
}

func Test_UpdateStock(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		stockRepo := &mocks.StockRepository{}
		stockRepo.On("UpdateStock", mock.Anything).Return(nil)

		uc := NewStockUsecase(stockRepo)

		err := uc.UpdateStock(&domain.Stock{})
		assert.NoError(t, err)
	})
}

func Test_DeleteStock(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		stockRepo := &mocks.StockRepository{}
		stockRepo.On("GetStockByID", mock.Anything).Return(&domain.Stock{}, nil)
		stockRepo.On("DeleteStock", mock.Anything).Return(nil)

		uc := NewStockUsecase(stockRepo)

		err := uc.DeleteStock(1)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		stockRepo := &mocks.StockRepository{}
		stockRepo.On("GetStockByID", mock.Anything).Return(&domain.Stock{}, errFoo)
		stockRepo.On("DeleteStock", mock.Anything).Return(nil)

		uc := NewStockUsecase(stockRepo)

		err := uc.DeleteStock(1)
		assert.Error(t, err)
	})
}
