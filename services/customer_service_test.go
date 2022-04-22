package services_test

import (
	"net/http"
	"testing"

	"github.com/parin/msisdn-country-validator/domain"
	"github.com/parin/msisdn-country-validator/services"
	"github.com/parin/msisdn-country-validator/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerWithError(t *testing.T) {

	t.Run("ok NewCustomerService", func(t *testing.T) {
		storeCalled := 0
		storeFunc := func() ([]*domain.Customer, *utils.ApplicationError) {
			storeCalled++
			return nil, nil
		}

		store := fakeStore{
			customeFetcher: storeFunc,
		}

		cs, err := services.NewCustomerService(store)
		assert.Nil(t, err)
		assert.NotNil(t, cs)

		_, ferr := cs.GetCustomers()
		assert.Nil(t, ferr)
		assert.True(t, storeCalled == 1)
	})

	t.Run("nok store GetCustomers error is propagated", func(t *testing.T) {
		storeCalled := 0
		storeFunc := func() ([]*domain.Customer, *utils.ApplicationError) {
			storeCalled++
			return nil, &utils.ApplicationError{
				Message:    "boom",
				StatusCode: http.StatusInternalServerError,
				Code:       "Fatal Error",
			}
		}

		store := fakeStore{
			customeFetcher: storeFunc,
		}

		cs, err := services.NewCustomerService(store)
		assert.Nil(t, err)
		assert.NotNil(t, cs)

		_, ferr := cs.GetCustomers()
		assert.NotNil(t, ferr)
		assert.True(t, storeCalled == 1)
		assert.True(t, ferr.Code == "Fatal Error")
	})

	t.Run("nok - NewCustomerService bad store returns error", func(t *testing.T) {
		cs, err := services.NewCustomerService(nil)
		assert.NotNil(t, err)
		assert.Nil(t, cs)
	})
}
type fakeStore struct {
	customeFetcher func() ([]*domain.Customer, *utils.ApplicationError)
}
func (f fakeStore) Close() error {
	return nil
}
func (f fakeStore) FetchCustomers() ([]*domain.Customer, *utils.ApplicationError) {
	return f.customeFetcher()
}
