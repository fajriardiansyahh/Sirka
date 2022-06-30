package resources

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	// _ "github.com/go-sql-driver/mysql" // An external package for handling mysql
	_ "github.com/lib/pq"
)

type CONNECTION_PARAM struct {
	LIFETIME string
	MAXOPEN  string
	MAXIDLE  string
}

// A global function to connect the application into the selected database on .env file
func Connect() *sql.DB {
	db, connString := setConnection()

	i, err := strconv.ParseInt(Load_Env("CONNECTION_LIFETIME"), 10, 64)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	max, err := strconv.ParseInt(Load_Env("CONNECTION_LIFETIME"), 10, 64)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	idle, err := strconv.ParseInt(Load_Env("CONNECTION_LIFETIME"), 10, 64)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	connection, err := sql.Open(db.DRIVER, connString)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	lifeTime := time.Duration(i)
	connection.SetConnMaxLifetime(time.Minute * lifeTime)
	connection.SetMaxOpenConns(int(max))
	connection.SetMaxIdleConns(int(idle))

	return connection
}

func setConnection() (DB_PARAM, string) {
	db := DB_PARAM{
		DRIVER:     Load_Env("DATABASE_DRIVER"),
		HOST:       Load_Env("DATABASE_URL"),
		PORT:       Load_Env("DATABASE_PORT"),
		USER:       Load_Env("DATABASE_USERNAME"),
		PASSWORD:   Load_Env("DATABASE_PASSWORD"),
		DB_NAME:    Load_Env("DATABASE_NAME"),
		CONNECTION: Load_Env("DATABASE_CONNECTION"),
		SSL:        Load_Env("DATABASE_SSLMODE"),
	}
	connectionString := ""

	switch db.DRIVER {
	case "mysql":
		connectionString = fmt.Sprintf(
			db.USER + ":@" +
				db.CONNECTION + "(" +
				db.HOST +
				db.PORT + ")/" +
				db.DB_NAME)

	case "oracle":
		connectionString = fmt.Sprintf(
			db.DRIVER + "://" +
				db.USER + ":" +
				db.PASSWORD + "@" +
				db.HOST + ":" +
				db.PORT + "/" +
				db.DB_NAME)

	case "postgres":
		connectionString = fmt.Sprintf(
			"host='" + db.HOST +
				"' port=" + db.PORT +
				" user=" + db.USER +
				" password='" + db.PASSWORD +
				"' dbname='" + db.DB_NAME +
				"' sslmode=" + db.SSL)
	}

	return db, connectionString
}
