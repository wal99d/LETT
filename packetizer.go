package LETT

import(
	//"fmt"
	//"errors"
	"encoding/gob"
	//"bytes"
	//"io/ioutil"
	"net"
)

//LETT Packet consists of [Header : Destination Address: Data itself of 2 Bytes in Size : Checksum to validate the bytes]

type LETTProtocol struct{
	Header uint8
	PacketNumber uint8
	SensorId uint8
	CollectorId uint8
    Readings uint16
	Stats uint8   //Used by COMMAND to set stats of addressed sensor using SensorId
    Checksum uint8
}

//These are const for Headers type in LETT protocol , maximum is 2^8 = 256 values
const (
	READINGS = 0x01
	COMMAND = 0x02
	READ_ACK = 0x03
	COMMAND_ACK = 0x04
	//Rest will be reserved for future use...
	
	//Stats 
	ON = 0x01
	OFF = 0x00
)

//This function will create brand new LETTPacket
func NewLETTPacket(conn *net.Conn, lettProto *LETTProtocol){
	err:= gob.NewEncoder(*conn).Encode(lettProto)
	if err != nil {
		panic(err)
	}
}

//This function will extract LETTPacket from Conn
func GetLETTProtocolPacket(conn *net.Conn) (*LETTProtocol , error){
	lettProto:= &LETTProtocol{}
	err:= gob.NewDecoder(*conn).Decode(lettProto)
	if err != nil {
		return nil, err
	}
	return lettProto , nil
}


