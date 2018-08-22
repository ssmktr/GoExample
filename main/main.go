package main

import (
	"net"
	"fmt"
	"github.com/unrolled/render"
)

var renderer render.Render

func onServer() {
	listener, err := net.Listen("tcp", ":2305")
	if err != nil {
		fmt.Errorf("Error listen : %v\n", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("Error Accept : %v\n")
			return
		}
		defer conn.Close()

		fmt.Printf("Connect remoteAddr : %v, localAddr : %v\n", conn.RemoteAddr().String(), conn.LocalAddr().String())

		go func (conn net.Conn) {
			data := make([]byte, 1024)

			for {
				n, err := conn.Read(data)
				if err != nil {
					fmt.Errorf("Error Read : %v\n", err)
					return
				}

				fmt.Println(string(data[:n]))

				_, err = conn.Write(data[:n])
				if err != nil {
					fmt.Errorf("Error Write : %v\n", err)
					return
				}

			}
		}(conn)
	}
}

func main() {
	onServer()
}



