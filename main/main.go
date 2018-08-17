package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/unrolled/render"
)

func server() {
	ln, err := net.Listen("tcp", ":2305")
	if err != nil {
		fmt.Printf("Error Listen : %v\n", err)
		return
	}
	defer ln.Close()

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error Accept : %v\n", err)
			continue
		}

		handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	defer c.Close()

	for {
		var msg string
		err := gob.NewDecoder(c).Decode(&msg)
		if err != nil {
			fmt.Printf("Error Decode : %v\n", err)
		} else {
			fmt.Printf("Received : %v\n", msg)
		}
	}
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:2305")
	if err != nil {
		fmt.Printf("Error Dial : %v\n", err)
		return
	}
	defer c.Close()

	for {
		var msg string
		fmt.Scanln(&msg)
		fmt.Println("Sending", msg)
		err = gob.NewEncoder(c).Encode(msg)
		if err != nil {
			fmt.Printf("Error Sending : %v\n", err)
			return
		}
	}
}

var renderer render.Render

func main() {
	// server()
	client()
}
