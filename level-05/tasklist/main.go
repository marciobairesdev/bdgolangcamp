package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

type Task struct {
	ID          string `json:"id" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type SimpleResponse struct {
	Message string `json:"message"`
}

func (i *Task) validate() error {
	errs := validate.Struct(*i)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		return err
	}

	return nil
}

var taskList []Task

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["description"]

	if ok && len(keys) == 1 && len(strings.TrimSpace(keys[0])) > 0 {
		filteredTasks := make([]Task, 0)
		for _, task := range taskList {
			if strings.Contains(strings.ToLower(task.Description), strings.ToLower(keys[0])) {
				filteredTasks = append(filteredTasks, task)
			}
		}

		json.NewEncoder(w).Encode(filteredTasks)
		return
	}

	json.NewEncoder(w).Encode(taskList)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if len(strings.TrimSpace(id)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: "Invalid id"})
		return
	}

	for _, task := range taskList {
		if task.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Task id %s not found", id)})
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload Task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: "Invalid payload"})
		return
	}

	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: err.Error()})
		return
	}

	if err = payload.validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: err.Error()})
		return
	}

	for _, task := range taskList {
		if task.ID == payload.ID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Task id %s already exists", payload.ID)})
			return
		}
	}

	taskList = append(taskList, payload)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payload)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var payload Task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: "Invalid payload"})
		return
	}

	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: err.Error()})
		return
	}

	if err = payload.validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: err.Error()})
		return
	}

	if payload.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: "Path id and payload id should be the same"})
		return
	}

	for i, task := range taskList {
		if task.ID == payload.ID {
			taskList[i] = payload
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(payload)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Task id %s not found", payload.ID)})
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if len(strings.TrimSpace(id)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: "Invalid id"})
		return
	}

	for i, task := range taskList {
		if task.ID == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Task id %s deleted", id)})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Task id %s not found", id)})
}

func main() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	taskList = make([]Task, 0)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/task", getAllTasks).Methods("GET")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	println("HTTP server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
