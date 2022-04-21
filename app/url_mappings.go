package app

import "github.com/parin/msisdn-country-validator/controllers"

func mapurls()  {
	router.GET("customers", controllers.GetCustomers)
	router.GET("countries", controllers.GetCountries)
}