package main

import (
	"be_good/src/db"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func get_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := db.GetData(10)
		// fmt.Println(data)
		ret, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ret)
	} else {

		w.Write([]byte("Error no such method available"))
	}
}

func main() {
	rout := mux.NewRouter()
	rout.HandleFunc("/get_data", get_data)
	serv := &http.Server{
		Handler:      rout,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(serv.ListenAndServe())
}
