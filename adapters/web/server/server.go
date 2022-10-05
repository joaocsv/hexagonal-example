package server

import (
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/joaocsv/hexagonal-example/adapters/web/handler"
	"github.com/joaocsv/hexagonal-example/app"
)

type WebServer struct {
	Service app.ProductServiceInterface
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w *WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)

	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           nil,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
