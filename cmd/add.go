/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the todoList",
	Long: `Using this function the user can new tasks to the todolist
	      The syntax of adding a new task is shown below
		[taskname] [priority] [deadline]`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			log.Fatal(err)
		}
		createTable := ` CREATE TABLE IF NOT EXISTS tasks(
		 id INTEGER PRIMARY KEY AUTOINCREMENT,
		 priority INTEGER DEFAULT 4,
		 title TEXT NOT NULL,
		 category TEXT,
		 creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		 check(priority > 0 AND priority < 5)		
		);`
		_, err = db.Exec(createTable)
		if err != nil {
			log.Fatal(err)
		}
		var Title string
		var priority int
		var category string

		fmt.Println("Please enter new task")
		fmt.Println("Please enter task title")
		fmt.Scan(&Title)
		fmt.Println("Please enter priority of task from 1-4")
		fmt.Scan(&priority)
		fmt.Println("Please enter category of task if any (e.g. college,work,chores etc.)")
		fmt.Scan(&category)

		addqry := "INSERT INTO tasks (priority,title,category) VALUES (?, ?, ?)"

		_, err = db.Exec(addqry, priority, Title, category)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Task Added Successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
