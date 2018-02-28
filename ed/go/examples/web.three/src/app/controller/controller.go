package controller

var (
	homeController home
)

func Startup() {
	homeController.registerRoutes()
}
