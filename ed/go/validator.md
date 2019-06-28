Validator
-

````sh
go get gopkg.in/go-playground/validator.v9
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
* unique # arrays & slices withoud duplicates
* alpha # alpha characters only
* alphanum
* numeric
* hexadecimal
* hexcolor
* rgb # rgb color string
* email
* url
* uri
* base64
* contains=@
* containsany=!@#?
* excludes=@
* excludesall=!@#?
* uuid
* uuid3
* uuid4
* uuid5
* ascii
* latitude
* longitude
* ip
* ipv4
* ipv6
* tcp_addr
* udp_addr
* hostname
* regexp=^[a-zA-Z0-9_]*$ (won't vork in v9)

Cross-Field Validation:

* eqfield # eqfield=ConfirmPassword
* nefield
* gtfield
* gtefield
* ltfield
* ltefield
* eqcsfield
* necsfield
* gtcsfield
* gtecsfield
* ltcsfield
* ltecsfield

````golang
type User struct {
  FirstName string `validate:"required"`
  LastName  string `validate:"-"` # don't validate
  Age       uint8  `validate:"gte=0,lte=130"`
  Email     string `validate:"required,email"`
}

validate := validator.New()
err := validate.Struct(user)
````

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

func validateKey(field validator.FieldLevel) bool {
  keyRegex := regexp.MustCompile("^[a-zA-Z0-9_\\-:]+$")

  return keyRegex.MatchString(field.Field().String())
}

func (cp ConfigPrototype) Validate() error {
  v := validator.New()
  v.RegisterValidation("key", validateKey)
  if err := v.Struct(cp); err != nil {
    return err
  }

  return nil
}
````
