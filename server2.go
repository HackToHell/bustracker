// server2
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type Message struct {
	utime int64
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

}
func opendatabase(busid int) {
	db, err := sql.Open("postgres", "user=postgres password='pol' dbname=test sslmode=disable ")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT * from bus_$1 ORDER BY utime DESC limit 1", busid)

	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
