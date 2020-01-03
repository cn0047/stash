package controller

var (
	homeController Home
	carsController Cars
	defaultController Default
)

func Startup() {
	homeController.registerRoutes()
	carsController.registerRoutes()
	defaultController.registerRoutes()
}
