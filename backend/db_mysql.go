package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

db, err := sql.Open("mysql", "user")

