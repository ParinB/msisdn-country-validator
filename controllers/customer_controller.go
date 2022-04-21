package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/parin/msisdn-country-validator/services"
	"net/http"
)

func GetCustomers(c *gin.Context)  {
	msisdn,err :=  services.CustomerService.GetCustomer()
	if err !=nil  {
			c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusOK,msisdn)
}