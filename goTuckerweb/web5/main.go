package main

import (
	"net/http"

	"github.com/JeonJe/goTuckerweb/web5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}

