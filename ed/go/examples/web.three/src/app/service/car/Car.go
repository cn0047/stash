package car

import (
	"app/service/car/dao"
	"app/service/car/request"
	"app/service/car/response"
)

func CreateNewCar(req request.New) response.Created {
	validateRequest(req)

	car := response.Created{}
	car.Id = 9
	car.Name = req.Name
	car.Vendor = req.Vendor

	dao.Save(car)

	return car
}

func validateRequest(req request.New) {
	//if req.Vendor != "BMW" {
	//	panic("")
	//}
}
