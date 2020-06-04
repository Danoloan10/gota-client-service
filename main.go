package main

import (
	"net/http"
	"sync"
	"log"
	"os"
    "os/signal"
    "syscall"
)

const CNN_HOSTNAME = "cnn.com"

func samplesHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// r.ParseForm()
		// guardar imagen en input/<sample_id>/<>.png
		// guardar campos en input/<sample_id>/data.json
		// añadir sample a una cola
	} else if r.Method == http.MethodGet {
		// mostrar cola
		// mostrar pendientes
		// ^ esto se puede hacer en Angular así que Redirect y ya
	}
	http.Redirect (w, r, "/", http.StatusFound)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func setupHTTPHandles(mux *http.ServeMux) {
	mux.HandleFunc("/samples", samplesHandle)
	mux.HandleFunc("/",        indexHandle)
}

func startUIServer(wg *sync.WaitGroup) *http.Server {
	mux := http.NewServeMux()
	srv := &http.Server{ Addr: "localhost:50505", Handler: mux }

	setupHTTPHandles(mux)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	} ()

	return srv
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	serviceWg := &sync.WaitGroup{}
	srv := startUIServer(serviceWg)
	// parar srv
	<-sigs
	srv.Close()
	serviceWg.Done()
}
