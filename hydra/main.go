package main

import (
	"fmt"
	"net/http"

	"github.com/loxt/mastering-go-programming/hydra/hlogger"
	"github.com/loxt/mastering-go-programming/hydra/shieldBuilder"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	builder := shieldBuilder.NewShieldBuilder()

	shield := builder.RaiseFront().RaiseBack().Build()
	logger.Printf("%+v \n", *shield)

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprint(w, "Welcome to the Hydra software system")

	logger.Println("Received an http Get request on root url")
}
