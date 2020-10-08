package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type test struct {
	Test int
	Test2 string
}

// GetUsers - getting list all users
func GetUsers (w http.ResponseWriter, r *http.Request) {

	data := test {
		Test: 1,
		Test2: "hey",
	}

	resp, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resp))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// CreateUser - creating a new user
func CreateUser (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

// GetUser - getting an user
func GetUser (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An users"))
}

// UpdateUser - updating an user
func UpdateUser (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an user"))
}

// DeleteUser - deleting an users
func DeleteUser (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete an user"))
}