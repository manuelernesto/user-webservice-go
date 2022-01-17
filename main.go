package main

import (
	"getstarted/demo/controllers"
	"net/http"
)

func main() {

	controllers.RegisterController()

	http.ListenAndServe(":3000", nil)

}
