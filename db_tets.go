// db_tets
package main

import (
	"code.google.com/p/goprotobuf/proto"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"msgproto"
	"time"
)

func main() {
	data := createmessage()
	fmt.Println(data)
	writedata(data)
	readata()
}

func createmessage() []byte {
	ProtoMessage := new(msgproto.Msg)
	ProtoMessage.Id = proto.Int32(12344)
	ProtoMessage.Lat = proto.Int64(12344)
	ProtoMessage.Long = proto.Int64(12344)
	ProtoMessage.Utime = proto.Int64(int64(time.Now().Unix()))
	data, err := proto.Marshal(ProtoMessage)
	if err != nil {
		fmt.Println(err)
	}
	return data

}

func readata() {
	db, err := sql.Open("postgres", "user=postgres password='pol' dbname=test sslmode=disable ")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT * from public.name")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var age int
		var id int
		if err := rows.Scan(&name, &age, &id); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s %d %d", name, age, id)
	}
	db.Close()
}
func writedata(data []byte) {
	ProtoMessage := new(msgproto.Msg)
	err := proto.Unmarshal(data, ProtoMessage)
	if err != nil {
	}
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
