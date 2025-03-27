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

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Display all the tasks",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		type taskrows struct {
			ID         int64  `field="id"`
			priority   int64  `field="priority"`
			title      string `field="title"`
			category   string `field="category"`
			createdate string `field="creation_date"`
		}
		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			log.Fatal(err)
		}
		var tasks *sql.Rows
		tasks, err = db.Query("SELECT * FROM tasks")
		if err != nil {
			log.Fatal(err)
		}
		defer tasks.Close()
		fmt.Println("ID Priority Title \t Category Creation Time")
		for tasks.Next() {
			list := new(taskrows)
			err = tasks.Scan(&list.ID, &list.priority, &list.title, &list.category, &list.createdate)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(list.ID)
			fmt.Print(" | ")
			fmt.Print(list.priority)
			fmt.Print("     | " + list.title + " | " + list.category + " | " + list.createdate + " |\n")

		}

	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
