// server
package main

import (
	"code.google.com/p/goprotobuf/proto"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"msgproto"
	"net"
	"os"
)

const (
	CONN_HOST = ""
	CONN_PORT = "8090"
	CONN_TYPE = "tcp"
)

func handleRequest(conn net.Conn) {
	data := make([]byte, 4096)
	// Read the incoming connection into the buffer.
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	protodata := new(msgproto.Msg)
	err = proto.Unmarshal(data[0:n], protodata)

	// Close the connection when you're done with it.
	conn.Close()
	go writedata(protodata)

}
func writedata(ProtoMessage *msgproto.Msg) {

	db, err := sql.Open("postgres", "user=postgres password='pol' dbname=test sslmode=disable ")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Exec("INSERT INTO public.data values($1,$2,$3,to_timestamp($4))", ProtoMessage.GetId(), ProtoMessage.GetLat(), ProtoMessage.GetLong(), ProtoMessage.GetUtime())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

}
func main() {
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}

}
