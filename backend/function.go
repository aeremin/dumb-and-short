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

type UrlDocument struct {
	Url string `firestore:"url"`
}

type CreateRequest struct {
	Url string `json:"url"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

func create(w http.ResponseWriter, r *http.Request) {
	var b CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Fprint(w, "Shit is broken!")
		return
	}

}

type RedirectRequest struct {
	Id string `json:"id"`
}

type RedirectResponse struct {
	Url string `json:"url"`
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var b RedirectRequest
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Fprintf(w, "Can't decode request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := client.Collection(urlsCollection).Doc(b.Id).Get(ctx)
	if err != nil {
		fmt.Fprintf(w, "Can't find a document: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var doc UrlDocument
	if err := d.DataTo(&doc); err != nil {
		fmt.Fprintf(w, "Can't parse a document: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := RedirectResponse{
		Url: doc.Url,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
