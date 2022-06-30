package resources

import (
	"database/sql"
	"log"
)

type DB_PARAM struct {
	DRIVER     string
	HOST       string
	PORT       string
	USER       string
	PASSWORD   string
	DB_NAME    string
	CONNECTION string
	SSL        string
}

type QueryResponse struct {
	Title  string
	Status int
	Data   struct{}
}

func Query(query string, args ...any) sql.Rows {
	connecting := Connect()
	defer connecting.Close()

	result, err := connecting.Query(query, args...)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return *result
}

func Read() {

}

func Update() {

}

func Delete() {

}
