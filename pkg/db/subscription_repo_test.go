package db_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/prodsub/pkg/db"
	uuid "github.com/satori/go.uuid"
	"github.com/tj/assert"
)

func TestSubscriptionRepo_Create(t *testing.T) {
	t.Run("CreateSubscription", func(t *testing.T) {
		productId := uuid.NewV4()
		subscr := db.Subscription{
			Price:     39.0,
			PlanName:  "Monthly",
			Duration:  800,
			ProductId: productId,
		}

		mock, gdb := prepare_db(t)

		mock.ExpectBegin()

		mock.ExpectQuery(`^INSERT INTO "subscriptions"*`).
			WithArgs(productId, subscr.PlanName, subscr.Duration, subscr.Price).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.NewV4()))

		mock.ExpectCommit()

		r := db.NewSubscriptionRepo(gdb)

		res, err := r.Create(&subscr)

		assert.NoError(t, err)
		assert.NotNil(t, res)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestSubscriptionRepo_Get(t *testing.T) {
	t.Run("SubscriptionFound", func(t *testing.T) {
		subId := uuid.NewV4()
		planName := "Monthly"

		mock, gdb := prepare_db(t)

		rows := sqlmock.NewRows([]string{"id", "plan_name"}).
			AddRow(subId, planName)

		mock.ExpectQuery(`^SELECT.*subscriptions.*`).
			WithArgs(subId, sqlmock.AnyArg()).
			WillReturnRows(rows)

		r := db.NewSubscriptionRepo(gdb)

		s, err := r.Get(subId)

		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.NotNil(t, s)

	})

}

func TestSubscriptionRepo_Update(t *testing.T) {

	t.Run("SubscriptionUpdatePlanName", func(t *testing.T) {
		subId := uuid.NewV4()
		newPlanName := "New Monthly Plan Name"

		mock, gdb := prepare_db(t)
		mock.ExpectBegin()

		mock.ExpectExec(`^UPDATE "subscriptions" SET*`).
			WithArgs(newPlanName, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		r := db.NewSubscriptionRepo(gdb)

		subscr, err := r.Update(subId, db.SubscriptionUpdateRequest{
			PlanName: newPlanName,
		})
		assert.NoError(t, err)
		assert.Equal(t, subId, subscr.Id)
	})
}

func TestSubscriptionRepo_List(t *testing.T) {

	t.Run("ListSubscriptionByProduct", func(t *testing.T) {
		productId := uuid.NewV4()
		subscr := db.Subscription{
			Id:       uuid.NewV4(),
			Price:    39.0,
			PlanName: "Monthly",
			Duration: 800,
		}

		mock, gdb := prepare_db(t)

		rows := sqlmock.NewRows([]string{"id", "plan_name", "price", "duration"}).
			AddRow(subscr.Id, subscr.PlanName, subscr.Price, subscr.Duration)

		mock.ExpectQuery(`^SELECT.*subscriptions.*`).
			WithArgs(productId).
			WillReturnRows(rows)

		r := db.NewSubscriptionRepo(gdb)

		list, err := r.List(productId)
		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.NoError(t, mock.ExpectationsWereMet())

	})
}

func TestSubscriptionRepo_Delete(t *testing.T) {
	t.Run("SubscriptionFound", func(t *testing.T) {
		subId := uuid.NewV4()

		mock, gdb := prepare_db(t)

		mock.ExpectBegin()

		mock.ExpectExec(`DELETE FROM "subscriptions"*`).
			WithArgs(subId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		r := db.NewSubscriptionRepo(gdb)

		err := r.Delete(subId)

		assert.NoError(t, err)
	})

	t.Run("SubscriptionNotFound", func(t *testing.T) {
		subId := uuid.NewV4()

		mock, gdb := prepare_db(t)

		mock.ExpectBegin()

		mock.ExpectExec(`DELETE FROM "subscriptions"*`).
			WithArgs(subId).
			WillReturnError(sql.ErrNoRows)

		r := db.NewSubscriptionRepo(gdb)

		err := r.Delete(subId)

		assert.Error(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)

	})
}
