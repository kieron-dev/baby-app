package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func main() {

	initDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		row := conn.QueryRow("select name from foo where id=$1", 1)
		var t string
		err := row.Scan(&t)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s: Hello, %s\n", t, r.UserAgent())
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func initDB() {
	connString := getConnString()
	var err error
	conn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

}

type vcapServices struct {
	Elephantsql []struct {
		Credentials struct {
			Uri string
		}
	}
}

func getConnString() string {
	s := os.Getenv("DB_CONN_STRING")
	if s != "" {
		return s
	}
	v := os.Getenv("VCAP_SERVICES")
	if v != "" {
		var services vcapServices
		err := json.Unmarshal([]byte(v), &services)
		if err != nil {
			log.Fatal(err)
		}
		return services.Elephantsql[0].Credentials.Uri
	}
	panic("could not find db creds")
}
