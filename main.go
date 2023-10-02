package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"bufio"
	"os"
)

func main() {
	args := os.Args
	db, err := sql.Open("sqlite3", "todos.db")
	// gormDB, _ := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if args[1] == "add" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Text to add: ")
		scanner.Scan()
		text := scanner.Text()
		query, _ := db.Prepare("INSERT INTO todos (text) VALUES (?)")
		_, err = query.Exec(text)
		if err != nil {
			fmt.Println(err)
		}
	} else if args[1] == "list" {
		query, err := db.Query("SELECT text FROM todos;")
		if err != nil {
			fmt.Println(err)
		}
		defer query.Close()

		for query.Next() {
			var text string
			err = query.Scan(&text)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(text)
		}
	} else if args[1] == "remove" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Text to remove: ")
		scanner.Scan()
		text := scanner.Text()
		query, _ := db.Prepare("DELETE FROM todos WHERE text=?")
		_, err = query.Exec(text)
		if err != nil {
			fmt.Println(err)
		}
	}
}