package services

import (
	"github.com/parin/msisdn-country-validator/domain"
	"github.com/parin/msisdn-country-validator/utils"
)

var (
	CountryService countryService
)
type countryService struct  {}

func (cs countryService) GetCountries() ([]domain.Country,*utils.ApplicationError)  {
	countries , err := domain.CountryDao.GetCountries()
	if err != nil {
		return nil,err
	}
	return  countries ,nil
}
