package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/parin/msisdn-country-validator/services"
	"net/http"
)

func GetCountries(c *gin.Context)   {
	countries,err :=  services.CountryService.GetCountries()
	if err !=nil  {
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusOK,countries)
}