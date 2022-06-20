package hdl

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><a href=/>Retour</a> Hello")
}
