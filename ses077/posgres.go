package main

import (
	"database/sql"
	"databases/sql"
	"fmt"

	"Githhub.com/lib/pq"
)

const (
	host="localhost"
	port=5432
	user="postgres"
	password=""
	dbname=""
)

var (db *sql.DB
err error)

func main(){
psqlinfo:= fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ",host,port,user,password,dbname)

db,err 
}

