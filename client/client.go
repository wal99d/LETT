package main 

import(
	lett "LETT"
	"fmt"
	"net"
)

func main(){

	lettproto:= &lett.LETTProtocol{
		Header: lett.COMMAND,
		PacketNumber: 1,
		SensorId: 1,
		Stats: lett.ON,
	}
	
	conn , err:= net.Dial("tcp", "127.0.0.1:23232")
	if err!=nil{
		fmt.Println(err)
	}
	//Send new LETT Protocol Packet
	lett.NewLETTPacket(&conn, lettproto)
	
	//Readout COMMANDs from LETT Server
	/*packet, err:=lett.GetLETTProtocolPacket(&conn)
	if err!=nil{
		fmt.Println(err)
	}
	switch(packet.Header){
		case READ_ACK:
			//in case we received READ_ACK then send next packet
			
	}*/
}
