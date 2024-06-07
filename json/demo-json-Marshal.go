package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(employee{ID: 1, EmployeeName: "John Doe", Tel: "1234567890", Email: "sarin@hotmail.com"})
	fmt.Println("JSON data: ", string(data))
}
