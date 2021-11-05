package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
    "github.com/rauf95/basic-crud.git"
)

type Sweet struct {
	ID		int	   `json:"id"`
	Title	string `json:"title"`
	Type	string `json:"type"`
	Kcal	int	   `json:"kcal"`
	Weight  string `json:"height"`
	Price	int    `json:"price"`
}


var Sweets = []Sweet {

	{
		ID: 1,
		Title: "Mochi",
		Type:"Japanese rice cake",
		Weight: "100 gr.",
		Kcal: 379,
		Price: 400,
	},
	{
		ID: 2,
		Title: "Baklava",
		Type:"traditional Turkish pastry",
		Weight: "100 gr.",
		Kcal: 433,
		Price: 583,
	},
	{
		ID: 3,
		Title: "Eclair",
		Type:"traditional French pastry",
		Weight: "100 gr.",
		Kcal: 414,
		Price: 325,
	},
	{
		ID: 4,
		Title: "Barmbrack",
		Type:"Sweet irish bread with raisin",
		Weight: "100 gr.",
		Kcal: 262,
		Price: 250,
	},
	{
		ID: 5,
		Title: "Cheesecake",
		Type:"international dessert, consisting of cream cheese and cookies",
		Weight: "100 gr.",
		Kcal: 276,
		Price: 320,
	},
}


func getSweets(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(Sweets)
	if err != nil{
		return
	}
}

func getSweet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params, _ := strconv.Atoi(mux.Vars(r)["id"])

	for _, item := range Sweets {
		if  item.ID == params{
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
		}
	}
}


func createSweets(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	var newSweet []Sweet

	if err != nil{}

	errSecond := json.Unmarshal(reqBody, &newSweet)
	if errSecond != nil {
		return
	}
	Sweets = append(Sweets, newSweet...)
}


func updateSweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params, _ := strconv.Atoi(mux.Vars(r)["id"])

	for index, item := range Sweets {
		if item.ID == params  {
			Sweets = append(Sweets[:index], Sweets[index+1])
			var sweet Sweet
			_ = json.NewDecoder(r.Body).Decode(&sweet)
			sweet.ID = params
			Sweets = append(Sweets, sweet)
			err := json.NewEncoder(w).Encode(sweet)
			if err != nil {
				return
			}
			return
		}
	}

}


func deleteSweets(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params, _ := strconv.Atoi(mux.Vars(r)["id"])

	for i, item := range Sweets {
		if item.ID == params{
			Sweets = append(Sweets[:i], Sweets[i+i])
			fmt.Printf("Sweet with ID %v is deleted", params)
		}
	}

}



func main(){


	router := mux.NewRouter()

	router.HandleFunc("/sweet", getSweets).Methods("GET")
	router.HandleFunc("/sweet/{id}", getSweet).Methods("GET")
	router.HandleFunc("/sweet", createSweets).Methods("POST")
	router.HandleFunc("/sweet/{id}", updateSweets).Methods("PUT")
	router.HandleFunc("/sweet/{id}", deleteSweets).Methods("DELETE")

	fmt.Printf("Start server 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}