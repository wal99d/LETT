package main

import (
	"log"
	lett "LETT"
	"net"
	"fmt"	
)
func main(){
	ln, err := net.Listen("tcp", ":23232")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	for {
		if conn, err := ln.Accept(); err == nil {
			//each client conn will be handled by single go routine
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	//He we need to handle each byte received from client and get all data need and display it
	buf:=make([]byte, 8)
	stringsRecevided:=[]string{""}
		    
	var lettproto *lett.LETTProtocol
	for {
		_,err:=conn.Read(buf)
		if err!=nil {
			//fmt.Println(err)
			break
		}
		lettproto,err=lett.GetLETTProtocolPacket(buf)
		if err!=nil{
			log.Println(err)
		}
		handleLettPacket(lettproto ,stringsRecevided)
	}
	//fmt.Println("Breaking For Loop")

	
}

func handleLettPacket(lettproto *lett.LETTProtocol , stringsRecevided []string){
	switch(lettproto.Header){
		case lett.COMMAND:
			fmt.Println("It's command Packet!")
			//TODO:Send the command to Sensor using SensorId
			break
		case lett.READINGS:
			//fmt.Println("It's Readings Packet!")
			//TODO:Save Readings into Mongodb with timestamp
			//fmt.Println(lettproto.PacketNumber)
			//if string(lettproto.Readings) != "\n"{
				stringsRecevided = append(stringsRecevided, string(lettproto.Readings))
				//}
			  fmt.Printf("%s", stringsRecevided)
			
			break
		case lett.READ_ACK:
			fmt.Println("It's Read_ack Packet!")
			break
		case lett.COMMAND_ACK:
			fmt.Println("It's Command_ack Packet!")
		default:
			//ignore this packet
			break
	}
}