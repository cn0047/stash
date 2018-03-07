package dao

import (
	"app/di"
	"app/service/car/response"
)

func Save(req response.Created) {
	di.GetMongoDB()
}
