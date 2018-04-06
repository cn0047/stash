Validator
-

````
go get gopkg.in/go-playground/validator.v9
````

````go
type User struct {
  FirstName string `validate:"required"`
  LastName  string `validate:"-"` # don't validate
  Age       uint8  `validate:"gte=0,lte=130"`
  Email     string `validate:"required,email"`
}

validate = validator.New()
err := validate.Struct(user)
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
* regexp=^[a-zA-Z0-9_]*$

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
