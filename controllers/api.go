package controllers

import (
	"Sirka/resources"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Display_User struct {
	Userid string `json:"userid"`
	Name   string `json:"name"`
}

func DisplayUser(r *http.Request) any {

	body := Display_User{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	query := "SELECT userid, name FROM users WHERE userid = $1"
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

		return response
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

	return result
}

func DisplayAllUsers(r *http.Request) any {
	query := "SELECT * FROM users"

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

		return response
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

	return result
}
