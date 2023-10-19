package database

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DBConnection
	product    *domain.Product
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	dialector := mysql.New(mysql.Config{
		Conn:       db,
		DriverName: "sqlmock",
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository = NewRepository(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Search() {
	s.mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(
		[]string{
			"sku", "name", "brand", "size", "price", "principal_image"}).AddRow(
		"test-sku", "name", "brand", "size", 100, "https://images"))

	res, err := s.repository.SearchProduct("test-sku")

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func (s *Suite) Test_repository_Search_Error() {
	s.mock.ExpectQuery("SELECT *").WillReturnError(errors.New("error not found"))

	res, err := s.repository.SearchProduct("test-sku")

	assert.Error(s.T(), err)
	assert.Nil(s.T(), res)
}

func (s *Suite) Test_repository_Get() {
	s.mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(
		[]string{
			"sku", "name", "brand", "size", "price", "principal_image"}).AddRow(
		"test-sku", "name", "brand", "size", 100, "https://images"))

	res, err := s.repository.GetProducts()

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func (s *Suite) Test_repository_Get_Error() {
	s.mock.ExpectQuery("SELECT *").WillReturnError(errors.New("error not found"))

	res, err := s.repository.GetProducts()

	assert.Error(s.T(), err)
	assert.Nil(s.T(), res)
}

func TestDBConnection_CreateProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := mysql.New(mysql.Config{
		Conn:       mockDb,
		DriverName: "mysql",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repository := NewRepository(db)

	err := repository.CreateProduct(*getProduct())

	assert.Nil(t, err)
}
