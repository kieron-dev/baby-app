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

type Hello struct {
	Name string `json:"name"`
}

func main() {

	initDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		row := conn.QueryRow("select name from foo where id=$1", 1)
		var t Hello
		err := row.Scan(&t.Name)
		if err != nil {
			log.Fatal(err)
		}

		out, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}

		origin := os.Getenv("CORS_ORIGIN")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		fmt.Fprint(w, string(out))
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
