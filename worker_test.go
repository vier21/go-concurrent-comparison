package DBT

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

type Data struct {
	Name   string `json: "firstName"`
	Email  string `json: "email"`
	Origin string `json: "origin"`
}

func NewData(name, email, origin string) *Data {
	return &Data{
		Name:   name,
		Email:  email,
		Origin: origin,
	}
}

func InsertMyJson() {

	var payload []Data
	
	newData := NewData("Samuel", "samuel@gmail.com", "Jakarta")
	for i := 0; i < 100000; i++ {
		
		payload = append(payload, *newData)
	}
	

	convJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("DBT/myjson.json", convJson, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func InsertSecJson() {

	var payload []Data

	newData := NewData("Samuel", "samuel@gmail.com", "Jakarta")

	
	for i := 0; i < 100000; i++ {
		
		payload = append(payload, *newData)
	}
	
	convJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}


	err = os.WriteFile("DBT/secjson.json", convJson, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func TestInsertMyJson(t *testing.T) {
	go InsertMyJson()
}
func TestInsertSecJson(t *testing.T) {
	InsertSecJson()
}
