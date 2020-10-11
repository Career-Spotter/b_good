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

// "GET :: /get_data?limit=some_int"
func get_data(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		keys, ok := r.URL.Query()["limit"]
		if !ok || len(keys[0]) < 1 {
			w.Write([]byte("missing limit param"))
			return
		}

		limit, err := strconv.ParseInt(keys[0], 10, 8)
		if err != nil {
			log.Fatal("couldn't convert to int")
		}
		data := db.GetData(int8(limit))
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
