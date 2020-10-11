package main

import (
	"net/http"

	"github.com/patilpankaj212/GoTest/simpleWebService/controller"
)

func main() {

	controller.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
