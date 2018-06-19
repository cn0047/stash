echo
-

````
var c echo.Context

// redirect
c.Redirect(http.StatusMovedPermanently, "<URL>")

// http://localhost/path/to/id/1
c.Param("id")

// http://localhost/path/to?id=1
c.QueryParam("id")

c.FormValue("name")

c.String(http.StatusOK, "Hello, World!")
c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
c.JSON(http.StatusOK, jsonStruct)
````
