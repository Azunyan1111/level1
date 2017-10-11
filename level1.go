package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/checkout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		d := json.NewDecoder(r.Body)
		var body interface{}
		e := d.Decode(&body)
		if e != nil {
			log.Fatalf("%+v\n", e)
		}
		fmt.Printf("%+v\n", body)
		fmt.Fprintln(w, "{}")
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
