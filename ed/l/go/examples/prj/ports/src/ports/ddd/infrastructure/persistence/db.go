package persistence

import (
	"fmt"

	"ports/ddd/domain/model"
)

var (
	db = map[string]model.PortEntity{}
)

func init() {
	db = make(map[string]model.PortEntity)
}

func Put(id string, m model.PortEntity) {
	db[id] = m
}

func Get(id string) (port model.PortEntity, err error) {
	p, in := db[id]
	if !in {
		return port, fmt.Errorf("port not found by id: %s", id)
	}

	return p, nil
}
