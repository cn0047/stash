package di

import (
	"github.com/app/products/dao"
)

// Init - initialize Dependency Injection Container.
func Init() (err error) {
	err = dao.Init()
	if err != nil {
		return
	}

	// Initialize other services, DAOs, etc...

	return
}
