package services

import (
	"github.com/parin/msisdn-country-validator/domain"
	"github.com/parin/msisdn-country-validator/utils"
)
var (CustomerService   customersService )

type  customersService struct {}

func (c customersService)  GetCustomer()  ([]*domain.Customer,*utils.ApplicationError) {
	var  validCustomers []*domain.Customer
	customers, err := domain.CustomerDao.FetchCustomers()
	if err != nil {
		return nil, err
	}
	for _ ,customer := range customers {
		customer.Validate()
		validCustomers = append(validCustomers,customer)
	}
	return validCustomers , nil
}

