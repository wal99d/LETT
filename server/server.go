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
	buf:=make([]byte, 1024)
	_,err:=conn.Read(buf)
	if err!=nil{
		fmt.Println(err)
	}
	lettproto,_:=lett.GetLETTProtocolPacket(buf[0:7])
	
	//fmt.Println(lettproto)
	handleLettPacketHeader(lettproto)
	//showDebugData(lettproto)
}

//This function is used for debuging purposes
func showDebugData(lettproto *lett.LETTProtocol){
	fmt.Println("LETT Header:",lettproto.Header)
	fmt.Println("LETT Readings:",lettproto.Readings)
}

func handleLettPacketHeader(lettproto *lett.LETTProtocol){
	switch(lettproto.Header){
		case lett.COMMAND:
			fmt.Println("It's command Packet!")
			//TODO:Send the command to Sensor using SensorId
			break
		case lett.READINGS:
			fmt.Println("It's Readings Packet!")
			//TODO:Save Readings into Mongodb with timestamp
			fmt.Println(lettproto.PacketNumber)
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
