package main

import (
	"basic-crud/controllers"
	"basic-crud/database"
	"basic-crud/entity"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main(){
	initDB()
	fmt.Printf("Start server 8080\n")

	router := mux.NewRouter().StrictSlash(true)
	Handlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Handlers(router *mux.Router) {
	router.HandleFunc("/sweet", controllers.GetSweets).Methods("GET")
	router.HandleFunc("/sweet/{id}", controllers.GetSweet).Methods("GET")
	router.HandleFunc("/sweet", controllers.CreateSweets).Methods("POST")
	router.HandleFunc("/sweet/{id}", controllers.UpdateSweets).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteSweets).Methods("DELETE")


}

func initDB()  {

	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "root",
			DB:         "learning_demo",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Sweet{})
}
