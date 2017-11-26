package main

import (
	"log"
	"net/http"

	"github.com/acoshift/gowebboard/app"
	"github.com/acoshift/gowebboard/ctrl"
)

func main() {
	mux := http.NewServeMux()

	userCtrl := ctrl.NewUserController()

	app.MountUserController(mux, userCtrl)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
