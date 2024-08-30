// Example for net.Conn
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	request := "GET / HTTP/1.0\r\n\r\n"
	conn.Write([]byte(request))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	fmt.Printf("Read %d bytes:\n%s\n", n, string(buf[:n]))
}
