package car

import (
	"app/service/car/request"
	"app/service/car/response"
)

func CreateNewCar(req request.New) response.Created {
	car := response.Created{}
	car.Id = 9
	car.Name = req.Name
	car.Vendor = req.Vendor

	return car
}
