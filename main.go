package main

import (
	"flag"
	 "github.com/parin/msisdn-country-validator/app"
	 "github.com/parin/msisdn-country-validator/domain"
	"github.com/parin/msisdn-country-validator/services"
	"log"
)

func main()  {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()
	db := flag.String("database", "./customer.db", "path to sqlite database")
	port := flag.String("port", ":5555", "port server listens on")
	flag.Parse()
	database, err := domain.NewCustomerDatabase(*db)
	if err != nil {
		return
	}
	cs,err := services.NewCustomerService(database)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	app := app.NewApp(database,cs)
	defer app.Close()
	app.SetUpRoutes()
	err = app.StartApplication(*port)
}