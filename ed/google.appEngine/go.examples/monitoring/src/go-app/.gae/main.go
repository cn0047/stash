package _gae

import (
	"google.golang.org/appengine"

	"go-app/route"
)

func init() {
	route.Page() // @important This route works by using prefix so it must be declared before else routes.
	route.Experiment()
	route.Home()
	route.Realtimelog()

	appengine.Main()
}
