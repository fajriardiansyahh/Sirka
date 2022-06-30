package resources

import (
	"encoding/json"
	"log"
	"net/http"
)

func APIHeaderJSON(w http.ResponseWriter, r *http.Request, maxBytes interface{}) {
	log.Println(r.URL.Path + r.URL.RawQuery)
	accept := r.Header.Get("Accept")
	contentType := r.Header.Get("Content-Type")
	if accept != "application/json" || contentType != "application/json" {
		msg := "Content-Type / Accept header is not application/json"
		log.Println(msg)
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		panic(msg)
	}

	if maxBytes != nil {
		switch v := maxBytes.(type) {
		case int64:
			log.Println(v)
			r.Body = http.MaxBytesReader(w, r.Body, maxBytes.(int64))
			r.ParseForm()
		}
	}

	w.Header().Set("Content-Type", "application/json")
}

func GetBody(r *http.Request, body interface{}) any {

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return body
}
