package helloworld

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("API", api)
}

type body struct {
	Id string `json:"id"`
}

func api(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
	return
}

func create(w http.ResponseWriter, r *http.Request) {
	var b body
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Fprint(w, "Shit is broken!")
		return
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var b body
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Fprint(w, "Shit is broken!")
		return
	}
}

var mux = newMux()

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/redirect", redirect)
	return mux
}
