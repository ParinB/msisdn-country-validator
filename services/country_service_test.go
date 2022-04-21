package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCountriesReturnsValues(t *testing.T) {
	countries,err := CountryService.GetCountries()
	assert.NotNil(t, countries)
	assert.Nil(t, err)
}