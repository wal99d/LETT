package main 

import(
	lett "LETT"
	"fmt"
	"net"
)

func main(){

	data := []byte(`{"Result":{ "Name":"Waleed" ,"Age":35}}`)
	var  checksum uint8 = 0
	for _,v :=range data{
		checksum ^=v
	}
	lettproto:= &lett.LETTProtocol{
		Header: lett.COMMAND,
		PacketNumber: 1,
		DestinationAddress: 1,
		Data: data,
		Checksum: checksum,
	}
	
	conn , err:= net.Dial("tcp", "127.0.0.1:23232")
	if err!=nil{
		fmt.Println(err)
	}
	
	lett.NewLETTPacket(&conn, lettproto)
}
