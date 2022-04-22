package app

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/parin/msisdn-country-validator/domain"
	"github.com/parin/msisdn-country-validator/services"
	"net/http"
)

//var (
//	router *gin.Engine
//)
//
//func init()  {
//	router = gin.Default()
//	router.Use(cors.Default())
//}
//
//func StartApp()  {
//	mapurls()
//	if err := router.Run(":3000"); err !=nil {
//		panic(err)
//	}
//}
type Application struct {
	database domain.Database
	customerService *services.CustomerService
	router *gin.Engine
}
//NewApp creates a new application connection
func NewApp(database domain.Database ,cs *services.CustomerService) *Application {
	router := gin.Default()
	router.Use(cors.Default())
	return &Application{
		database: database,
		customerService: cs,
		router: router,
	}
}
// Set Up Application Routes
func (a Application) SetUpRoutes() {
	//TODO. implement on another component for scalability
	a.router.GET("customers",a.GetCustomersController)
}
//GetCustomersHandler handles get customers request
func (a Application) GetCustomersController(c *gin.Context) {
	customers, err := a.customerService.GetCustomers()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, customers)
}
//Close  terminates all dependencies
func (a Application) Close() error {
	return a.database.Close()
}
// StartApplication launches app on a certain port
func (a Application) StartApplication(port string) error {
	if err := a.router.Run(port); err != nil {
		return fmt.Errorf("unable to start application: %w", err)
	}
	return nil
}
