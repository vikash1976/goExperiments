package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
    "strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()

	for {
		con, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go serve(con)

	}
}

func serve(con net.Conn) {
	defer con.Close()
	//io.WriteString(con, "I see you connected")
	scanner := bufio.NewScanner(con)

    var i int
    var rMethod, rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		//fmt.Println(ln)
        //fmt.Fprintf(con, "i heard you say: %s", ln)
		//io.WriteString(con, ln)
        if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			//io.WriteString(con, "Bye")
			break
		}
        i++
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	io.WriteString(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/plain\r\n")
	io.WriteString(con, "\r\n")
	io.WriteString(con, body)

	fmt.Println("Code got here")

}
