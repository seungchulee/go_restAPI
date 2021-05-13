package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ACCESS!!")
	fmt.Fprintf(w, "Welcome home!")
}

func create(w http.ResponseWriter, r *http.Request) {
	var Event event
	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(respBody, &Event)
	events = append(events, Event)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(events)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, event := range events {
		if event.ID == id {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(event)
		}
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updateEvent event
	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR!")
	}
	json.Unmarshal(respBody, &updateEvent)
	fmt.Println(updateEvent)

	for i, event := range events {
		if event.ID == id {
			event.Description = updateEvent.Description
			event.Title = updateEvent.Title
			events = append(events[:i], event)
			json.NewEncoder(w).Encode(event)
		}
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/get/{id}", getOne).Methods("GET")
	router.HandleFunc("/getAll", getAll).Methods("GET")
	router.HandleFunc("/update/{id}", update).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", router))
}
