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
			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	//He we need to handle each byte received from client and get all data need and display it
	lettproto,_:=lett.GetLETTProtocolPacket(&conn)
	fmt.Println("LETT Haader:",lettproto.Header) 
}
