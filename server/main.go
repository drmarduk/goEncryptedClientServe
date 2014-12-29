// server

package main

import (
	"encoding/json"
	"log"
	"net"
)

var (
	Room Channel
)

func main() {
	Room = Channel{Name: "Test", Topic: "dööörp", Users: []User{}}
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
		go handleMsg(conn)
	}

}

func handleMsg(c net.Conn) {
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		log.Println(err.Error())
		return
	}
	buf = buf[:n]

	//msg := convertBuffer(buf)
	var cmd Command
	err = json.Unmarshal( /*msg*/ buf, &cmd)
	if err != nil {
		log.Println(err.Error())
	}

	switch cmd.Command {
	case "JOIN":
		if cmd.User != "" {
			Room.AddUser(cmd.User, cmd.PublicKey, c)
			log.Printf("%s added to %s\n", cmd.User, Room.Name)
		} else {
			log.Println("Could not join empty user.")
		}
	}

	c.Write([]byte("dörp\n"))
}

// testable :)
func convertBuffer(b []byte) string {
	return string(b)
}
