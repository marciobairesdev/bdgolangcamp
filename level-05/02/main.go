/*
 * 		- Statement:
 * 			Create a RESTful API implementing all the CRUD Actions, and integrate using gorilla/mux.
 * 			Create a Product struct with the following attributes: [ID string, Code string, Name string, Price float64].
 * 			Create a ProductInventory struct with the following attributes: [Product Product, Quantity int].
 * 			Then to simulate a table in memory, create a var inventory as a []ProductInventory.
 * 			Now, Create the functions [Add, Update, Delete, Get] that will be executed against inventory.
 * 			For example: Add will add a new ProductInventory into the inventory ([]ProductInventory).
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

type SimpleResponse struct {
	Message string `json:"message"`
}

type Product struct {
	ID    string  `json:"id" validate:"required"`
	Code  string  `json:"code" validate:"required"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"gt=0"`
}

type ProductInventory struct {
	Product  Product `json:"product" validate:"required"`
	Quantity int     `json:"quantity" validate:"gt=0"`
}

func (i *ProductInventory) validate() error {
	errs := validate.Struct(*i)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		return err
	}

	return nil
}

var productInvetory []ProductInventory

func getAllProductsFromInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productInvetory)
}

func getProductFromInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if len(strings.TrimSpace(id)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: "Invalid id"})
		return
	}

	for _, pi := range productInvetory {
		if pi.Product.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(pi)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Product id %s not found", id)})
}

func addProductToInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload ProductInventory
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

	for _, pi := range productInvetory {
		if pi.Product.ID == payload.Product.ID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Product id %s already exists", payload.Product.ID)})
			return
		}
	}

	productInvetory = append(productInvetory, payload)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payload)
}

func updateProductInInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload ProductInventory
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

	for i, pi := range productInvetory {
		if pi.Product.ID == payload.Product.ID {
			productInvetory[i] = payload
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(payload)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Product id %s not found", payload.Product.ID)})
}

func deleteProductFromInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if len(strings.TrimSpace(id)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&SimpleResponse{Message: "Invalid id"})
		return
	}

	for i, pi := range productInvetory {
		if pi.Product.ID == id {
			productInvetory = append(productInvetory[:i], productInvetory[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Product id %s deleted", id)})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&SimpleResponse{Message: fmt.Sprintf("Product id %s not found", id)})
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

	productInvetory = make([]ProductInventory, 0)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/inventory", getAllProductsFromInventory).Methods("GET")
	router.HandleFunc("/inventory/{id}", getProductFromInventory).Methods("GET")
	router.HandleFunc("/inventory", addProductToInventory).Methods("POST")
	router.HandleFunc("/inventory", updateProductInInventory).Methods("PUT")
	router.HandleFunc("/inventory/{id}", deleteProductFromInventory).Methods("DELETE")

	println("HTTP server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
