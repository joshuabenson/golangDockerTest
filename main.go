package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(ctx context.Context) []*User {

	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/josh_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	result, err := db.ExecContext(ctx,
		"INSERT INTO users (name) VALUES ($1)",
		"gopher",
		27,
	)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*User
	for results.Next() {
		var u User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}
func addUser() []*User {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/josh_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("INSERT INTO users (`name`) VALUES () ")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*User
	for results.Next() {
		var u User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage Bitch!")
	fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Endpoint Hit: /users")
	json.NewEncoder(w).Encode(users)
}

func handleAddUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: /addUser")
	//json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	http.HandleFunc("/addUser", handleAddUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
