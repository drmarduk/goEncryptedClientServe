// server

package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		// process message
		go handleInputConnection(conn)
	}

}

func handleInputConnection(c net.Conn) {
	buf := make([]byte, 1024)
	_, err := c.Read(buf)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Debug: Input from Connection: %v\n", buf)
	c.Write([]byte("dörp\n"))
}
