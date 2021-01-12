package gorm

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnMySQL() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "sunhapark:root@tcp(127.0.0.1:3306)/")
	defer db.Close() // Close the DB
	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to : ", version)
	//(Outputs) Connected to :  10.5.8-MariaDB
}
