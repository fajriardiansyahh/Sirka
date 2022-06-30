package controllers

import (
	"Sirka/resources"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Display_User struct {
	Userid string `json:"Userid"`
	Name   string `json:"Name"`
}

func DisplayUser(w http.ResponseWriter, r *http.Request) {
	resources.APIHeaderJSON(w, r, 1048576)

	body := Display_User{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	query := "SELECT Userid, Name FROM Users WHERE Userid=?"
	user := resources.Query(query, body.Userid)
	var result []Display_User

	defer user.Close()
	if user.Next() == false {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			panic(err)
		}

		response := map[string]any{
			"Title":       "StatusNotFound",
			"Status":      http.StatusNotFound,
			"URL":         r.Host + r.URL.Path,
			"User-Device": hostname,
			"Method":      r.Method,
		}
		jsonResp, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResp)
		// return jsonResp
	}

	for user.Next() {
		var display_user = Display_User{}
		err := user.Scan(&display_user.Userid, &display_user.Name)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		result = append(result, display_user)
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	w.Write(jsonResp)
	// return jsonResp
}

func DisplayAllUsers(w http.ResponseWriter, r *http.Request) {
	resources.APIHeaderJSON(w, r, 1048576)
	query := "SELECT * FROM Users"

	user := resources.Query(query)
	var result []Display_User

	defer user.Close()
	if user.Next() == false {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			panic(err)
		}

		response := map[string]any{
			"Title":       "StatusNotFound",
			"Status":      http.StatusNotFound,
			"URL":         r.Host + r.URL.Path,
			"User-Device": hostname,
			"Method":      r.Method,
		}
		jsonResp, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResp)
		// return jsonResp
	}

	for user.Next() {
		var display_user = Display_User{}
		err := user.Scan(&display_user.Userid, &display_user.Name)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		result = append(result, display_user)
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	w.Write(jsonResp)
	// return jsonResp
}
