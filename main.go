package main

import (
	app "go-jwt/app"
	"net/http"
)

func main() {
	http.Handle("/api", app.ValidateJWT(app.HomeHandler))
	http.HandleFunc("/jwt", app.GetJwtHandler)

	http.ListenAndServe(":3500", nil)

}
