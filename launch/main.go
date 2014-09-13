package main

import (
	"code.google.com/p/gorest"
	"github.com/jabong/register"
	"net/http"
)

func main() {
	gorest.RegisterService(&register.RegisterService{})
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8787", nil)
}
