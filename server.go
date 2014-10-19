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
	//	"strconv"
)

const (
	CONN_HOST = "192.168.1.101"
	CONN_PORT = "8090"
	CONN_TYPE = "tcp"
)

//func createdatabd(n int) {
//	db, err := sql.Open("postgres", "user=postgres password='pol' dbname=test sslmode=disable ")
//	var doesitexist string
//	if err != nil {
//		fmt.Println(err)
//	}

//	for i := 0; i <= n; i++ {
//		rows, err := db.Exec("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_name=$1)", n).Scan(&doesitexist)
//		fmt.Println(rows)
//		if err != nil {
//			fmt.Println(err)
//		}
//		if doesitexist == "t" {

//		} else {
//			rows, err := db.Exec("CREATE TABLE $1(Lat numberic(10),Long numeric(10),utime timestamp", n)
//			if err != nil {
//				fmt.Println(err)
//			}
//		}

//	}
//}

func handleRequest(conn net.Conn) {
	data := make([]byte, 4096)
	// Read the incoming connection into the buffer.
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	protodata := new(msgproto.Msg)
	err = proto.Unmarshal(data[0:n], protodata)
	fmt.Println(protodata.GetId())
	fmt.Println(protodata.GetLat())

	// Close the connection when you're done with it.
	conn.Close()
	//go writedata(protodata)

}
func writedata(ProtoMessage *msgproto.Msg) {

	db, err := sql.Open("postgres", "user=postgres password='pol' dbname=test sslmode=disable ")
	if err != nil {
		fmt.Println(err)
	}
	//var query string
	//string = "INSERT INTO bus_" + string(ProtoMessage.GetId()) + " values(" + strconv.FormatInt(ProtoMessage.GetLat(), 10) + "," + strconv.FormatInt(ProtoMessage.GetLong(), 10) + "," + strconv.FormatInt(ProtoMessage.GetUtime(), 10)
	rows, err := db.Exec("INSERT INTO $1 values($2,$3,to_timestamp($4))", ProtoMessage.GetId(), ProtoMessage.GetLat(), ProtoMessage.GetLong(), ProtoMessage.GetUtime())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

}
func main() {
	//createdatabd(10)
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Hi")
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
