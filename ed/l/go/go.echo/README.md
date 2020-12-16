echo
-

[doc](https://echo.labstack.com/guide)

Examples:
[customContext](https://echo.labstack.com/guide/context)
[templates](https://echo.labstack.com/guide/templates)
[dbUserLogin](https://echo.labstack.com/cookbook/twitter)

````golang
var c echo.Context

c.Request().RequestURI
c.Request().Referer()

// http://localhost/path/to/id/1
c.Param("id")

// http://localhost/path/to?id=1
c.QueryParam("id")

c.FormValue("name")

// JSON in POST
input := VO{
    Email string `json:"email" param:"email" query:"email" form:"email"`
}
err := ctx.Bind(&input)

// redirect
c.Redirect(http.StatusFound, "<URL>")
c.Response().Header().Set(echo.HeaderLocation, "img")

c.String(http.StatusOK, "Hello, World!")
c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
c.JSON(http.StatusOK, jsonStruct)
````
