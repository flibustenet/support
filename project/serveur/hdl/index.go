package hdl

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>Index <a href=hello>Hello</a></html>")
}
