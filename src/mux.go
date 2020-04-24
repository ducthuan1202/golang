package main

import (
	"fmt"
	"net/http"
	"routes"
)

func main() {

	routes.AppRoute.Initial()

	http.Handle("/", routes.AppRoute.RootRoute)

	fmt.Println("Server Go started at <http://localhost:3000>")

	http.ListenAndServe(":3000", nil)
}
