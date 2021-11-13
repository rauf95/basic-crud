package database

import (
	"basic-crud/entity"
	"gorm.io/gorm"
	"log"
)


var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)

	if err != nil{
		return err
	}
	log.Println("success connection")
	return nil
}
func Migrate(table *entity.Sweet) {
	err := Connector.AutoMigrate(&table)
	if err != nil {
		return
	}
	log.Println("Table migrated")
}



