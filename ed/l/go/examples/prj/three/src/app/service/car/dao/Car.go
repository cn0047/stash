package dao

import (
	"fmt"
	"time"

	"app/di"
	"app/service/car/response"
)

func Save(req response.Created) {
	di.GetMongoDB()
}

func Delete(id int) {
	time.Sleep(5 * time.Second)
	fmt.Printf("⏱ Deleted.\n")
	panic("RUNTIME-ERROR-CAR-DELETE-1: Some exception.")
}
