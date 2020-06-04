package ui

import (
	"os"
	"net/http"
	"sync"
	"log"
	"mime/multipart"
	"strings"
	"io"

	"github.com/satori/go.uuid"
)

const (
	gotaRoot    = "gota" // TODO hacer que dependa de una envariable, e independiente del so
	gotaSamples = gotaRoot + "/samples"
)

func saveFile(ifile multipart.File, head *multipart.FileHeader) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		// TODO
		return
	}
	dir  := gotaSamples + "/" + id.String()
	path := dir + "/" + head.Filename
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Println(err)
		// TODO
		return
	}
	ofile, err := os.Create(path)
	if err != nil {
		log.Println(err)
		//TODO
		return
	} else {
		defer ofile.Close()
	}
	io.Copy(ofile, ifile)
}

func addSample(r *http.Request) {
	// cogemos fichero del form
	if err := r.ParseForm(); err != nil {
		// TODO notificar error a IU
		log.Println(err)
		return
	}
	file, head, err := r.FormFile("file")
	if err != nil {
		// TODO notificar error a IU
		log.Println(err)
		return
	} else {
		defer file.Close()
	}

	// guardar imagen

	contentType := head.Header.Get("Content-Type")
	contentType  = strings.Split(contentType, "/")[0]

	// comprobación de seguridad
	if contentType == "image" {
		saveFile(file, head)
	} else {
		// TODO notificar error a IU
		log.Println("not an image")
	}

	// guardar campos en input/<sample_id>/data.json
	// añadir sample a una cola
}

func samplesHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		addSample(r)
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

func StartUIServer(wg *sync.WaitGroup) *http.Server {
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

