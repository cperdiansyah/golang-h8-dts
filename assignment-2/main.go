package main

import (
	"github.com/cperdiansyah/golang-h8-dts/assignment-2/app"
	"github.com/cperdiansyah/golang-h8-dts/assignment-2/controller"
	"github.com/cperdiansyah/golang-h8-dts/assignment-2/repository"
	"github.com/cperdiansyah/golang-h8-dts/assignment-2/service"
)

func main() {
	db := app.NewDB()
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	router := app.OrderRoute(orderController)

	router.Run()
}
