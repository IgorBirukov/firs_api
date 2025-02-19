package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//type task struct {
//Task   string `json:"task"`
//IsDone bool   `json:"is_done"`
//}

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

		n0 := Task{Task: ""}
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

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		mvars := mux.Vars(r)
		id := mvars["id"]
		var task Task

		err := DB.First(&task, id).Error
		if err != nil {
			http.Error(w, "Tasks not found", http.StatusNotFound)
			return
		}

		var patch map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
			http.Error(w, "Errors of json", http.StatusBadRequest)
		}

		DB.Model(&task).Updates(patch)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		mvars := mux.Vars(r)
		id := mvars["id"]
		var task Task

		err := DB.First(&task, id).Error
		if err != nil {
			http.Error(w, "Tasks not found", http.StatusNotFound)
			return
		}
		fmt.Println(task)

		if err := DB.Delete(&task).Error; err != nil {
			http.Error(w, "Deleted error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func main() {

	InitDB()
	DB.AutoMigrate(&Task{})
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
