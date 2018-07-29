package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	
	"google.golang.org/appengine"
)

var people []Person

// our main function
func main() {

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Boeing", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	people = append(people, Person{ID: "4", Firstname: "Kosta", Lastname: "Fotiadis"})
	people = append(people, Person{ID: "5", Firstname: "Cam", Lastname: "Day"})
	people = append(people, Person{ID: "6", Firstname: "Mark", Lastname: "Briggs", Address: &Address{City: "Woking", State: "Hampshire"}})
	people = append(people, Person{ID: "7", Firstname: "Pogo", Lastname: "Poggers", Address: &Address{City: "Transylvania", State: "Noon State"}})
	people = append(people, Person{ID: "8", Firstname: "Lulu", Lastname: "Mcfarlane", Address: &Address{City: "Kruekham", State: "Bront"}})
	people = append(people, Person{ID: "9", Firstname: "Kendall", Lastname: "Atkins", Address: &Address{City: "Midrord", State: "Zriglore"}})
	people = append(people, Person{ID: "10", Firstname: "Heath", Lastname: "Norton", Address: &Address{City: "Oawrapolis", State: "Cria"}})
	people = append(people, Person{ID: "11", Firstname: "Hammond", Lastname: "Erickson", Address: &Address{City: "Drita", State: "Pliplens"}})
	people = append(people, Person{ID: "12", Firstname: "Fatimah", Lastname: "Vazquez", Address: &Address{City: "Ofrosburgh", State: "Vlaaphora"}})


	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	//log.Fatal(http.ListenAndServe(":8080", router))

	http.Handle("/", router)
	appengine.Main()

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
