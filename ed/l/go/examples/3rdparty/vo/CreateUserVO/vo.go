package CreateUserVO

import (
	"errors"
)

type Instance struct {
	// errors contains all validation errors.
	errors map[string]string

	name  string // intentionally non-exported
	email string // intentionally non-exported
}

// New creates new value object instance and performs validation.
func New(data map[string]string) (Instance, error) { // Instance intentionally value, not pointer
	vo := Instance{}
	vo.errors = make(map[string]string)

	vo.initName(data)
	vo.initEmail(data)

	if len(vo.errors) > 0 {
		return vo, errors.New("invalid VO")
	}

	return vo, nil
}

// initName performs naive name validation (just for sake of example)
// and provides 100% guarantee that vo contains valid name.
func (i *Instance) initName(data map[string]string) {
	name, exists := data["name"]
	if !exists || name == "" {
		i.errors["name"] = "invalid name"
		return
	}

	i.name = name
}

// initEmail same like initName.
func (i *Instance) initEmail(data map[string]string) {
	email, exists := data["email"]
	if !exists || email == "" {
		i.errors["email"] = "invalid email"
		return
	}

	i.email = email
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
