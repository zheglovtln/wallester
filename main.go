package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/gin-gonic/gin"
	"wallester/controllers"
	"wallester/db"
	"wallester/models"
	"wallester/repository"
	"wallester/routes"
	websockets "wallester/websocket"
)

func main() {
	//Connect to db
	var db = db.Connect()
	//Migration
	db.AutoMigrate(&models.Customer{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	//Repository
	var customerRepository = repository.NewCustomerRepository(db)
	var customerController = controllers.New(customerRepository)
	router := gin.Default()
	//Websockets
	dispatcher := websockets.NewDispatcher(customerRepository)
	go dispatcher.Run()
	wsHandler := websockets.NewHandlers(dispatcher)

	staticBox := rice.MustFindBox("static")
	router.StaticFS("/static", staticBox.HTTPBox())
	router.HTMLRender = gorice.New(rice.MustFindBox("/views"))

	routes.Attach(router, customerController)
	websockets.Attach(router,wsHandler)

	router.Run(":80")


}