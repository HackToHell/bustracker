// client
package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"math/rand"
	"msgproto"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Start")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go send_data(create_message(), wg)
	}
	wg.Wait()
}
func create_message() []byte {

	ProtoMessage := new(msgproto.Msg)
	ProtoMessage.Id = proto.Int32(int32(rand.Intn(1000)))
	ProtoMessage.Lat = proto.Int64(int64(rand.Intn(1000)))
	ProtoMessage.Long = proto.Int64(int64(rand.Intn(1000)))
	ProtoMessage.Utime = proto.Int64((time.Now().Unix()))
	data, err := proto.Marshal(ProtoMessage)
	if err != nil {
		fmt.Println(err)
	}
	return data

}

func send_data(data []byte, wg sync.WaitGroup) {
	fmt.Println("Data")
	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println(err)
	}
	n, err := conn.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	wg.Done()
}
