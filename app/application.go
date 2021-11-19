package app

import (
	"github.com/Lozovoi-Rodion/bookstore_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
