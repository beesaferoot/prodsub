package db_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/prodsub/pkg/db"
	uuid "github.com/satori/go.uuid"
	"github.com/tj/assert"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestProductRepo_Create(t *testing.T) {
	t.Run("CreateProduct", func(t *testing.T) {
		product := db.Product{
			Price:            700,
			Name:             "Shoe",
			Type:             db.PhysicalProduct,
			ProductAttribute: datatypes.JSON([]byte(`{"weight": 0,"dimensions": "10"}`)),
		}

		mock, gdb := prepare_db(t)

		mock.ExpectBegin()

		mock.ExpectQuery(`INSERT INTO*`).
			WithArgs(product.Name, product.Type, sqlmock.AnyArg(), product.Price,
				sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.NewV4()))

		mock.ExpectCommit()

		r := db.NewProductRepo(gdb)

		res, err := r.Create(&product)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestProductRepo_Get(t *testing.T) {
	t.Run("ProductFound", func(t *testing.T) {
		var productId = uuid.NewV4()
		var name = "shoe"

		mock, gdb := prepare_db(t)

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(productId, name)

		mock.ExpectQuery(`^SELECT.*products.*`).
			WithArgs(productId, sqlmock.AnyArg()).
			WillReturnRows(rows)

		r := db.NewProductRepo(gdb)

		p, err := r.Get(productId)

		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
		assert.NotNil(t, p)

	})
}

func TestProductRepo_List(t *testing.T) {

	t.Run("ListAll", func(t *testing.T) {
		p := &db.Product{
			Id:    uuid.NewV4(),
			Name:  "Shoe",
			Price: 80,
			Type:  db.DigitalProduct,
		}

		mock, gdb := prepare_db(t)

		rows := sqlmock.NewRows([]string{"name", "price", "type"}).
			AddRow(p.Name, p.Price, p.Type)

		mock.ExpectQuery(`^SELECT.*products.*`).
			WithArgs().
			WillReturnRows(rows)

		r := db.NewProductRepo(gdb)

		list, err := r.List("")
		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("ListByType", func(t *testing.T) {
		p := &db.Product{
			Id:    uuid.NewV4(),
			Name:  "Shoe",
			Price: 80,
			Type:  db.DigitalProduct,
		}

		mock, gdb := prepare_db(t)

		rows := sqlmock.NewRows([]string{"name", "price", "type"}).
			AddRow(p.Name, p.Price, p.Type)

		mock.ExpectQuery(`^SELECT.*products.*`).
			WithArgs(fmt.Sprintf("%d", p.Type)).
			WillReturnRows(rows)

		r := db.NewProductRepo(gdb)

		list, err := r.List(fmt.Sprintf("%d", p.Type))
		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestProductRepo_Delete(t *testing.T) {
	t.Run("ProductFound", func(t *testing.T) {
		productId := uuid.NewV4()

		mock, gdb := prepare_db(t)

		mock.ExpectBegin()

		mock.ExpectExec(`DELETE FROM "products"*`).
			WithArgs(productId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		r := db.NewProductRepo(gdb)

		err := r.Delete(productId)

		assert.NoError(t, err)
	})

	t.Run("ProductNotFound", func(t *testing.T) {

		productId := uuid.NewV4()

		mock, gdb := prepare_db(t)

		mock.ExpectBegin()

		mock.ExpectExec(`^DELETE FROM "products"*`).
			WithArgs(productId).
			WillReturnError(sql.ErrNoRows)

		r := db.NewProductRepo(gdb)

		err := r.Delete(productId)

		assert.Error(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func prepare_db(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {

	db, mock, err := sqlmock.New() // mock sql.DB
	assert.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gdb, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	return mock, gdb
}
