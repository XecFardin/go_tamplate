package app

import (
	"net/http"

	"github.com/XecFardin/go_tamplate/controller"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
