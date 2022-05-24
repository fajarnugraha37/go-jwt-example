package app

import (
	"fmt"
	"net/http"
)

func GetJwtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header["jwt"] != nil {
		if r.Header["jwt"][0] != api_key {
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}
