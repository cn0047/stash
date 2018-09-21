package experiment

import (
	"net/http"
	"fmt"
	"runtime"
)

func RuntimeNumCPUHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NumCPU: %v\n", runtime.NumCPU())
}
