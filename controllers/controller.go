package controllers

import (
	"encoding/json"
	"basic-crud/database"
	"basic-crud/entity"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetSweets(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(entity.Sweets)
	if err != nil{
		return
	}
}

func GetSweet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params, _ := strconv.Atoi(mux.Vars(r)["id"])

	for _, item := range entity.Sweets {
		if  item.ID == params{
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
		}
		database.Connector.First(entity.Sweets, params)
	}
}


func CreateSweets(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	var newSweet []entity.Sweet

	if err != nil{}

	errSecond := json.Unmarshal(reqBody, &newSweet)
	if errSecond != nil {
		return
	}
	entity.Sweets = append(entity.Sweets, newSweet...)

	database.Connector.Create(&entity.Sweets)
}


func UpdateSweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params, _ := strconv.Atoi(mux.Vars(r)["id"])

	for index, item := range entity.Sweets {
		if item.ID == params  {
			entity.Sweets = append(entity.Sweets[:index], entity.Sweets[index+1])
			var sweet entity.Sweet
			_ = json.NewDecoder(r.Body).Decode(&sweet)
			sweet.ID = params
			entity.Sweets = append(entity.Sweets, sweet)
			err := json.NewEncoder(w).Encode(sweet)
			if err != nil {
				return
			}
			return
		}
		database.Connector.Save(&entity.Sweets)
	}
}


func DeleteSweets(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params, _ := strconv.Atoi(mux.Vars(r)["id"])

	for i, item := range entity.Sweets {
		if item.ID == params{
			entity.Sweets = append(entity.Sweets[:i], entity.Sweets[i+i])
			fmt.Printf("Sweet with ID %v is deleted", params)
		}
		database.Connector.Where("id = ?", params).Delete(&entity.Sweets)
	}

}


