GIN
-

[docs](https://github.com/gin-gonic/gin)

````
view - ok
controller - cool
model - no
db - no
````

````go
type CreateRequest struct {
    Username string `json:"username" form:"username"`
    Phone    string `json:"phone" form:"phone"`
    Password string `json:"password" form:"password"`
}
````
