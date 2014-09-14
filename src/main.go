package main

import (
	"code.google.com/p/gorest"
	"webservice"
	"net/http"
)

func main() {
	gorest.RegisterService(&webservice.RegisterService{})
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8787", nil)
}
