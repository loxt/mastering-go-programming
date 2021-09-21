package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)	
}

func sroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Hydra software system")
}