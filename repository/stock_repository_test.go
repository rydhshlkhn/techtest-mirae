package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rydhshlkhn/techtest-mirae/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func Test_CrerateStock(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	stockRepo := NewStockRepository(db)

	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	SQL := "INSERT INTO \"stocks\" (.+) VALUES (.+)"
	mock.ExpectBegin()
	mock.ExpectQuery(SQL).WillReturnRows(addRow)
	mock.ExpectCommit()
	var stock domain.Stock
	stockRepo.CreateStock(&stock)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func Test_GetAllStock(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	stockRepo := NewStockRepository(db)

	stocks := sqlmock.NewRows([]string{"id", "name", "code", "price", "fequency", "volume"}).
		AddRow(1, "Mahaka Media Tbk.", "ABBA", 1000, 1, 100)

	SQL := "SELECT * FROM \"stocks\""
	mock.ExpectQuery(regexp.QuoteMeta(SQL)).WillReturnRows(stocks)
	res := stockRepo.GetAllStock()
	assert.NotNil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_GetStockByID(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	stockRepo := NewStockRepository(db)

	stocks := sqlmock.NewRows([]string{"id", "name", "code", "price", "fequency", "volume"}).
		AddRow(1, "Mahaka Media Tbk.", "ABBA", 1000, 1, 100)

	SQL := "SELECT (.+) FROM \"stocks\" WHERE id =(.+)"
	mock.ExpectQuery(SQL).WillReturnRows(stocks)
	_, res := stockRepo.GetStockByID(1)
	assert.Nil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func Test_UpdateStock(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	stockRepo := NewStockRepository(db)

	SQL := "UPDATE \"stocks\" SET .+"
	mock.ExpectBegin()
	mock.ExpectExec(SQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := stockRepo.UpdateStock(&domain.Stock{})
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func Test_DeleteStock(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	stockRepo := NewStockRepository(db)

	SQL := "DELETE FROM \"stocks\" WHERE id = .+"
	mock.ExpectBegin()
	mock.ExpectExec(SQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := stockRepo.DeleteStock(1)
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}
