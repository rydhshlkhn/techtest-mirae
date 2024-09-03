package resthandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rydhshlkhn/techtest-mirae/domain"
	mocks "github.com/rydhshlkhn/techtest-mirae/mocks/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, reqBody, reqParam string
	wantUsecaseError        error
	wantRespCode            int
}

var (
	errFoo = errors.New("something error")
)

func TestRestHandler_createStock(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"name":"Mahaka Media Tbk.","code":"AAAA","price":500,"frequency":1,"volume":1200}`, wantUsecaseError: nil, wantRespCode: http.StatusCreated,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"name":"Mahaka Media Tbk.","code":"AAAA","price":500,"frequency":1,"volume":1200}`, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"name":"Mahaka Media Tbk.","code":"AAAA","price":err,"frequency":1,"volume":1200}`, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			stockUsecase := new(mocks.StockUsecase)
			NewStockHandler(app, stockUsecase)
			stockUsecase.On("CreateStock", mock.Anything).Return(tt.wantUsecaseError)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/stock", strings.NewReader(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)

			assert.Equal(t, tt.wantRespCode, res.StatusCode)
		})
	}
}

func TestRestHandler_getAllStock(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			stockUsecase := new(mocks.StockUsecase)
			NewStockHandler(app, stockUsecase)
			stockUsecase.On("GetAllStock", mock.Anything).Return([]domain.Stock{}, tt.wantUsecaseError)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/stock", strings.NewReader(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)

			assert.Equal(t, tt.wantRespCode, res.StatusCode)
		})
	}
}

func TestRestHandler_getStockByID(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqParam: "/1", wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqParam: "/1", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", reqParam: "/err", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			stockUsecase := new(mocks.StockUsecase)
			NewStockHandler(app, stockUsecase)
			stockUsecase.On("GetStockByID", mock.Anything).Return(&domain.Stock{}, tt.wantUsecaseError)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/stock"+tt.reqParam, strings.NewReader(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)

			assert.Equal(t, tt.wantRespCode, res.StatusCode)
		})
	}
}

func TestRestHandler_updateStock(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqParam: "/1", reqBody: `{"name":"Mahaka Media Tbk.","code":"AAAA","price":500,"frequency":1,"volume":1200}`, wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqParam: "/err", reqBody: `{"name":"Mahaka Media Tbk.","code":"AAAA","price":500,"frequency":1,"volume":1200}`, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", reqParam: "/1", reqBody: `{"name":"Mahaka Media Tbk.","code":"AAAA","price":500,"frequency":1,"volume":1200}`, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #4: Negative", reqParam: "/1", reqBody: `{"name":"Mahaka Media Tbk.","code":"AAAA","price":err,"frequency":1,"volume":1200}`, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			stockUsecase := new(mocks.StockUsecase)
			NewStockHandler(app, stockUsecase)
			stockUsecase.On("UpdateStock", mock.Anything).Return(tt.wantUsecaseError)

			req := httptest.NewRequest(http.MethodPut, "/api/v1/stock"+tt.reqParam, strings.NewReader(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)

			assert.Equal(t, tt.wantRespCode, res.StatusCode)
		})
	}
}

func TestRestHandler_deleteStock(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqParam: "/1", wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqParam: "/1", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", reqParam: "/err", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			stockUsecase := new(mocks.StockUsecase)
			NewStockHandler(app, stockUsecase)
			stockUsecase.On("DeleteStock", mock.Anything).Return(tt.wantUsecaseError)

			req := httptest.NewRequest(http.MethodDelete, "/api/v1/stock"+tt.reqParam, strings.NewReader(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req)

			assert.Equal(t, tt.wantRespCode, res.StatusCode)
		})
	}
}
