package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//var renderer render.Render
//
//func onServer() {
//	listener, err := net.Listen("tcp", ":2305")
//	if err != nil {
//		fmt.Errorf("Error listen : %v\n", err)
//		return
//	}
//	defer listener.Close()
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Errorf("Error Accept : %v\n")
//			return
//		}
//		defer conn.Close()
//
//		fmt.Printf("Connect remoteAddr : %v, localAddr : %v\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
//
//		go func (conn net.Conn) {
//			data := make([]byte, 1024)
//
//			for {
//				n, err := conn.Read(data)
//				if err != nil {
//					fmt.Errorf("Error Read : %v\n", err)
//					return
//				}
//
//				fmt.Println(string(data[:n]))
//
//				_, err = conn.Write(data[:n])
//				if err != nil {
//					fmt.Errorf("Error Write : %v\n", err)
//					return
//				}
//
//			}
//		}(conn)
//	}
//}

func main() {
	//onServer()
	testMysql()
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

	go conneded(conn)
}

func conneded(conn *sql.DB) {
	for {
		conn.Query("select 1")

		time.Sleep(120)
	}
}



