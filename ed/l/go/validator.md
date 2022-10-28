Validator
-

[docs](https://github.com/go-playground/validator)

````sh
go get gopkg.in/go-playground/validator.v9
go get github.com/go-playground/validator/v10
````

* len=10
* max=10
* min=1
* eq=10
* ne=10
* oneof=red green
* oneof=5 7 9
* gt=10
* lt=10
* unique (arrays & slices withoud duplicates)
* alpha (alpha characters only)
* alphanum
* numeric
* rgb (rgb color string)
* contains=@
* containsany=!@#?
* excludes=@
* excludesall=!@#?
* regexp=`^[a-zA-Z0-9_]*$` (won't vork in v9)

Cross-Field Validation:

* eqfield # eqfield=ConfirmPassword
* nefield

````golang
type User struct {
  FirstName string `validate:"required"`
  LastName  string `validate:"-"`              // don't validate
  Age       uint8  `validate:"gte=0,lte=130"`
  Email     string `validate:"required,email"`
}

validate := validator.New()
err := validate.Struct(user)
````

Register custom validation:

````golang
package service

import (
  "gopkg.in/go-playground/validator.v9"
  "regexp"
)

type ConfigPrototype struct {
  Key         string      `json:"key" validate:"required,key"`
  Description string      `json:"description" validate:"max=1000"`
  Type        string      `json:"type" validate:"required,oneof=int float32 string bool object array"`
  Value       interface{} `json:"value" validate:"required"`
}

func (cp ConfigPrototype) Validate() error {
  v := validator.New()
  v.RegisterValidation("key", validateKey)
  if err := v.Struct(cp); err != nil {
    return err
  }

  return nil
}

var keyRegex = regexp.MustCompile("^[a-zA-Z0-9_\\-:]+$")

func validateKey(field validator.FieldLevel) bool {
  return keyRegex.MatchString(field.Field().String())
}

var (
  // ResourceNameRegex - represents regex pattern to validate resource name.
  ResourceNameRegex = regexp.MustCompile(`^([\w:-])+$`)
)

func IsResourceNameValid(value string) bool {
  return ResourceNameRegex.MatchString(value)
}

````
