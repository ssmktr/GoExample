package main

import (
	"GoExample/gametabledata"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"GoExample/httpservermanager"
	)

func main() {
	gametabledata.RunGameTableDataServer()
	httpservermanager.RunHttpServer()
	//tcpservermanager.OnServer()
	//testMysql()
}

func testMysql() {
	fmt.Println("===== mysql start =====")
	conn, err := sql.Open("mysql", "root:ball2305@tcp(localhost:3306)/db_test")
	if err != nil {
		fmt.Printf("Error mysql open : %v\n", err)
		return
	}
	defer func() {
		conn.Close()
		fmt.Println("===== mysql finish =====")
	}()

	fmt.Println("Success mysql open")

	rows, err := conn.Query("select uid, id, pw from tb_test")
	if err != nil {
		fmt.Printf("Error select tables : %v\n", err)
		return
	}

	for rows.Next() {
		var uid int
		var id string
		var pw string
		err := rows.Scan(&uid, &id, &pw)
		if err != nil {
			fmt.Printf("Error Scan : %v\n", err)
			return
		}

		fmt.Printf("uid : %v, id : %v, pw : %v\n", uid, id, pw)
	}
}



