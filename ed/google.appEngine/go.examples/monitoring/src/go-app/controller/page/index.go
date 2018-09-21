package page

import (
	"net/http"
)

var (
	dir = http.Dir("../.gae/template/page/")
	IndexHandler = http.StripPrefix("/page/", http.FileServer(dir))
)