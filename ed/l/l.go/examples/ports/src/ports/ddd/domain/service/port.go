package service

import (
	"ports/ddd/domain/model"
	"ports/ddd/infrastructure/persistence"
)

func PutPort(port model.PortEntity) error {
	persistence.Put(port.ID, port)

	return nil
}

func GetPort(id string) (model.PortEntity, error) {
	return persistence.Get(id)
}
