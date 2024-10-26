package helloworld

import (
  "encoding/json"
  "fmt"
  "html"
  "net/http"

  "github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
   functions.HTTP("HelloHTTP", helloHTTP)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func helloHTTP(w http.ResponseWriter, r *http.Request) {
  mux.ServeHTTP(w, r)
  return
  var d struct {
    Name string `json:"name"`
  }
  if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
    fmt.Fprint(w, "Shit is broken!")
    return
  }
  if d.Name == "" {
    fmt.Fprint(w, "Hello, anonymous!")
    return
  }
  fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}


func hello(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w,"Hello World!")
}

func login(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w,"Login from /subroute/login")
}

var mux = newMux()

func newMux() *http.ServeMux {
   mux := http.NewServeMux()
   mux.HandleFunc("/hello", hello)
   mux.HandleFunc("/subroute/login", login)
   return mux
}
