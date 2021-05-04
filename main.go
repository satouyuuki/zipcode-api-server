package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"server2/pkg/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", show)
	// api.FetchIndex()
	panic(http.ListenAndServe(":8080", r))
}

func show(w http.ResponseWriter, r *http.Request) {
	data := api.FetchIndex()
	// json形式に変換
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(bytes)))
}