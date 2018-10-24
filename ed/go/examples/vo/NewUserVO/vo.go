package NewUserVO

import (
	"errors"
)

type Instance struct {
	errors map[string]string

	name  string
	email string
}

func New(data map[string]string) (Instance, error) {
	vo := Instance{}
	vo.errors = make(map[string]string)

	vo.initName(data)
	vo.initEmail(data)

	if len(vo.errors) > 0 {
		return vo, errors.New("Invalid VO.")
	}

	return vo, nil
}

func (i *Instance) initName(data map[string]string) {
	name, exists := data["name"]
	if !exists {
		i.errors["name"] = "name is required."
		return
	}
	if name == "" {
		i.errors["name"] = "name cannot be blank."
		return
	}

	i.name = name
}

func (i *Instance) initEmail(data map[string]string) {
	name, exists := data["email"]
	if !exists {
		i.errors["email"] = "email is required."
		return
	}
	if name == "" {
		i.errors["email"] = "email cannot be blank."
		return
	}

	i.name = name
}

func (i Instance) GetErrors() map[string]string {
	return i.errors
}

func (i Instance) GetName() string {
	return i.name
}

func (i Instance) GetEmail() string {
	return i.email
}
