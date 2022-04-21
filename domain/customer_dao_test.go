package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCustomerReturnsError(t *testing.T) {
	customers,err := CustomerDao.FetchCustomers()
	assert.NotNil(t, err)
	assert.Nil(t, customers)
}