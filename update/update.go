package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("AAAAAA", 1)

	stmt2, _ := db.Prepare("delete from  usuarios where id = ?")
	stmt2.Exec(1)
}
