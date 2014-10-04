// db_tets
package msgproto

import (
	"code.google.com/p/goprotobuf/proto"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"msgproto"
	"time"
)

func main() {
	readmessage(createmessage())
}

func createmessage() []byte {
	ProtoMessage := new(msgproto.Msg)
	ProtoMessage.Id = proto.Int32(12344)
	ProtoMessage.Lat = proto.Int64(12344)
	ProtoMessage.Long = proto.Int64(12344)
	ProtoMessage.Utime = proto.Int64(int64(time.Now().Unix()))
	proto.Marshal(ProtoMessage)

}

func readmessage(data []byte) {
	ProtoMessage := new(msgproto.Msg)
	err = proto.Unmarshal(data[0:n], ProtoMessage)
	if err != nil {
	}
	fmt.Println(ProtoMessage.getId())

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
