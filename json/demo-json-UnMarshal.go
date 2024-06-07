// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type employee struct {
// 	ID           int
// 	EmployeeName string
// 	Tel          string
// 	Email        string
// }

// func main() {

// 	e := employee{}

// 	err := json.Unmarshal([]byte(`{"ID":1,"EmployeeName":"John Doe","Tel":"1234567890","Email":"asljdhaslhf"}`), &e)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(e.Email)

// }
