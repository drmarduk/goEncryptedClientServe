// client
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/otr"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:1337")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "hi")
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	log.Println(resp)
}

func encryptOtr(msg string) (string, error) {
	conversation := otr.Conversation{}

}
