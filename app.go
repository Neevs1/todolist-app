/*
 * A simple to-do list app using Go programming language, Cobra library for CLI and SQLite for database
 * The aim is to create a cross platform (windows,linux,mac) app
 * Author : Neevs1 (github.com/Neevs1)
 *
 */

package main

import (
	_ "database/sql"
	"fmt"
	_ "fmt"
	_ "strconv"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/spf13/cobra"
)

func main() {
	fmt.Println("Welcome to this app!")

}
