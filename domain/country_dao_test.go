package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesReturnsValues(t *testing.T) {
	countries,err := CountryDao.GetCountries()
	assert.NotNil(t, countries)
	assert.Nil(t, err)
}