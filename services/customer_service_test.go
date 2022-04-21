package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//returns error since  customer db cannot be found
func TestGetCustomerWithError(t *testing.T) {
	customer,err := CustomerService.GetCustomer()
	assert.NotNil(t, err)
	assert.Nil(t, customer)
}

//func TestGetCustomerNoError(t *testing.T) {
//	customer,err := CustomerService.GetCustomer()
//	assert.NotNil(t, err)
//	assert.Nil(t, customer)
//}