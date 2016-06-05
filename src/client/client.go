package main 

import(
	lett "LETT"
	"fmt"
	"net"
)

func main(){
	
	conn , err:= net.Dial("tcp", "127.0.0.1:23232")
	if err!=nil{
		fmt.Println(err)
	}
	data:="Thermostat:13C,door:opend\n"
	rdng:= []byte(data)
	/*var rdngSize int
	if len(rdng)<=8 {
		rdngSize= 1
	}else{
		rdngSize= len(rdng)/8
	}*/
	for i:=0; i<len(rdng); i++{
		lettproto:= &lett.LETTProtocol{
			Header: lett.READINGS,
			PacketNumber: uint8(i),
			SensorId: 1,
			CollectorId:1,
			Readings: rdng[i],
		}
		//Marshal new LETT Protocol Packet
		packet:=lettproto.Marshal()
		conn.Write(packet)
		fmt.Printf("Sending...Packet %d\n",i)
	}
	conn.Close()

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