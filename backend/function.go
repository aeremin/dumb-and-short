package helloworld

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"log"
	"net/http"
	"strconv"
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
	client, err = firestore.NewClientWithDatabase(ctx, firestore.DetectProjectID, "dumb-and-short")
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	functions.HTTP("API", api)
}

func handleCors(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return false
}

type MetadataDocument struct {
	NextId int `firestore:"next_id"`
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
	if handleCors(w, r) {
		return
	}
	var b CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Fprintf(w, "Can't decode request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := CreateResponse{}

	err := client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		metadataDocRef := client.Collection(urlsCollection).Doc("metadata")
		metaSnap, err := tx.Get(metadataDocRef)
		if err != nil {
			return err
		}

		metadata := MetadataDocument{}
		if err := metaSnap.DataTo(&metadata); err != nil {
			return err
		}

		doc := UrlDocument{
			Url: b.Url,
		}

		if err := tx.Create(client.Collection(urlsCollection).Doc(strconv.Itoa(metadata.NextId)), doc); err != nil {
			return err
		}

		metadata.NextId++
		if err := tx.Set(metadataDocRef, metadata); err != nil {
			return err
		}

		resp.Id = strconv.Itoa(metadata.NextId - 1)

		return nil
	})

	if err != nil {
		fmt.Fprintf(w, "Internal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Fprintf(w, "Can't encode a response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

type RedirectRequest struct {
	Id string `json:"id"`
}

type RedirectResponse struct {
	Url string `json:"url"`
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if handleCors(w, r) {
		return
	}
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

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Fprintf(w, "Can't encode a response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}
