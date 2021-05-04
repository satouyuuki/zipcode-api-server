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
	r.HandleFunc("/api/{zipcode:[0-9]+}", findAddress)
	// api.FetchIndex()
	panic(http.ListenAndServe(":8080", r))
}

func jsonEncode(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}

func show(w http.ResponseWriter, r *http.Request) {
	// json形式に変換
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonEncode(api.FetchIndex())))
}

func findAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["zipcode"]
	data := api.FetchByKey(id)
	if len(data) == 0 {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonEncode(data)))
}

type NotFound404 struct {
	Message string `json:"message"`
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		message := NotFound404{"ソースが見つかりませんでした"}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsonEncode(message)))
	}
}