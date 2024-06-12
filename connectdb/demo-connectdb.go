package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func creatingTable(db *sql.DB) {
	query := `CREATE TABLE users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		create_at DATETIME
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func Insert(db *sql.DB) {
	var username, password string
	fmt.Scan(&username)
	fmt.Scan(&password)
	create_at := time.Now()
	result, err := db.Exec("INSERT INTO users (username, password, create_at) VALUES (?, ?, ?)", username, password, create_at)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	fmt.Println("Inserted record with ID:", id)

}

func delete(db *sql.DB) {
	var id int
	fmt.Scan(&id)
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Deleted record with ID:", rowsAffected)

}

// func query(db *sql.DB) {
// 	type course struct {
// 		id         int
// 		coursename string
// 		price      float64
// 		instructor string
// 	}

// 	var c course
// 	for {
// 		var inputID int
// 		fmt.Scan(&inputID)
// 		query := "SELECT id, coursename, price, instruuctor FROM onlinecourse WHERE id = ?"
// 		if err := db.QueryRow(query, inputID).Scan(&c.id, &c.coursename, &c.price, &c.instructor); err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(c.id, c.coursename, c.price, c.instructor)
// 	}
// }

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.100.81:3304)/coursedb")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to the database")
	}
	//query(db)
	//creatingTable(db)
	//Insert(db)
	delete(db)

}
