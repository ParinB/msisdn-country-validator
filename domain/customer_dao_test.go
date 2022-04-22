package domain_test

import (
	"testing"

	"github.com/parin/msisdn-country-validator/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomer(t *testing.T) {
	t.Run("ok >>>>> fetch customers", func(t *testing.T) {
		db, err := domain.NewCustomerDatabase("testdb/customer.db")
		assert.True(t, err == nil)
		t.Cleanup(func() {
			db.Close()
		})

		customers, derr := db.FetchCustomers()
		assert.Nil(t, derr)
		assert.True(t, len(customers) > 0)
	})

	t.Run("nok >>>>>>>>> error getting a non existing database", func(t *testing.T) {
		db, err := domain.NewCustomerDatabase("testdb/database-does-not-exist")
		assert.Nil(t, err)

		customers, derr := db.FetchCustomers()
		assert.NotNil(t, derr)
		assert.True(t, len(customers) == 0)
	})
}
