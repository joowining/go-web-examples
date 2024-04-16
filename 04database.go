package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:korean2078@(127.0.0.1:3306)/web_ide?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	er := db.Ping()
	if er != nil {
		log.Fatalln(er)
	}

	//	query := `
	//		CREATE TABLE users (
	//			id INT AUTO_INCREMENT,
	//			username TEXT NOT NULL,
	//			password TEXT NOT NULL,
	//			created_at DATETIME,
	//			PRIMARY KEY (id)
	//		)
	//	`
	//	_, err = db.Exec(query)

	//	username := "johndoe"
	//	password := "secret"
	//	created_at := time.Now()

	//	result, err := db.Exec(`INSERT INTO users ( username, password, created_at ) VALUES (?,?,?)`, username, password, created_at)

	//	userID, err := result.LastInsertId()
	//	fmt.Println(userID)

	/*	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)*/

	//	query := `SELECT id, username, password, created_at FROM users WHERE id =?`
	//	err = db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	//	fmt.Println(id, username, password, createdAt)

	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		users = append(users, u)
	}
	err = rows.Err()
	fmt.Println(users)

}
