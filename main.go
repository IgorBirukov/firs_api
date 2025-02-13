package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

var mTask string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	w.Header().Set("Content-Type", "application/json")
	var result []struct {
		Task string
	}

	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
	fmt.Println(result)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		n0 := task{Task: ""}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&n0)
		if err != nil {
			panic(err)
		}
		mTask = n0.Task
		DB.Create(&n0)
		json.NewEncoder(w).Encode(&n0)
		fmt.Println("Получено значени:", n0.Task)

	}

}

func main() {

	InitDB()
	DB.AutoMigrate(&task{})
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
