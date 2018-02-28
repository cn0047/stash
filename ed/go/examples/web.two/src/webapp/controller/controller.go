package controller

var (
	homeController home
	shopController shop
)

func Startup() {
	homeController.registerRoutes()
	shopController.registerRoutes()
}
