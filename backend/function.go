package helloworld

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var client *firestore.Client
var ctx = context.Background()

const urlsCollection = "urls"

var mux = newMux()

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/redirect", redirect)
	return mux
}

func api(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

func init() {
	var err error
	client, err = firestore.NewClient(ctx, "alice-larp")
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	functions.HTTP("API", api)
}

type body struct {
	Id string `json:"id"`
}

type UrlDocument struct {
	Url string `firestore:"url"`
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

	d, err := client.Collection(urlsCollection).Doc(b.Id).Get(ctx)
	if err != nil {
		fmt.Fprint(w, "Shit is broken 2!")
		return
	}
	var doc UrlDocument
	if err := d.DataTo(&doc); err != nil {
		fmt.Fprint(w, "Shit is broken 3!")
	}

	http.Redirect(w, r, doc.Url, 301)
}
