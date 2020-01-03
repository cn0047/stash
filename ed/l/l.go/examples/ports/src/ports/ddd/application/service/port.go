package service

import (
	"fmt"
	"os"

	"ports/ddd/domain/model"
	"ports/ddd/domain/service"
	"ports/ddd/infrastructure/file_reader"
)

func LoadPortsFromJSONFile(file string, cb func(port model.PortEntity) error) error {
	var err error

	// todo: discuss validation and check whether file exists.
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file for init DB, error: %#v", err)
	}

	err = file_reader.Read(f, cb)
	if err != nil {
		return fmt.Errorf("failed read data from file, error: %#v", err)
	}

	return nil
}

func PutPort(port model.PortEntity) error {
	return service.PutPort(port)
}

func GetPort(id string) (model.PortEntity, error) {
	return service.GetPort(id)
}
