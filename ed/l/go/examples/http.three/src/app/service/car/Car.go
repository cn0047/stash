package car

import (
	"strconv"

	"app/service/car/dao"
	"app/service/car/request"
	"app/service/car/response"
)

func CreateNewCar(req request.New) response.Created {
	validateRequestCreateNewCar(req)

	car := response.Created{}
	car.Id = 9
	car.Name = req.Name
	car.Vendor = req.Vendor

	dao.Save(car)

	return car
}

func validateRequestCreateNewCar(req request.New) {
	if req.Vendor != "BMW" {
		panic("Vendor must be BMW.")
	}
}

func DeleteCarById(idString string) response.Deleted {
	validateRequestDeleteCarById(idString)
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic("RUNTIME-ERROR-CAR-1: " + err.Error())
	}

	car := response.Deleted{}
	car.Id = id

	go dao.Delete(id)

	return car
}

func validateRequestDeleteCarById(id string) {
	// ctype digit
}
