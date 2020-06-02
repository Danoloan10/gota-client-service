package main

import "net/http"
import "strings"

func data(w http.ResponseWriter, r *http.Request) {
	if strings.Split(r.RemoteAddr, ":")[0] == "127.0.0.1" {
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
	} else {
		// TODO conexion externa
		http.Error (w, "q haces", http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/data", data)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	http.ListenAndServe(":50505", nil)
}
