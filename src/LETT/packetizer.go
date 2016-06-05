package LETT

import(
	//"fmt"
	//"errors"
	//"encoding/gob"
	"encoding/binary"
	"bytes"
	//"bufio"
	//"io/ioutil"
	//"net"
)

//LETT Packet consists of [Header : Destination Address: Data itself of 2 Bytes in Size : Checksum to validate the bytes]

type LETTProtocol struct{
	Header uint8
	PacketNumber uint8
	SensorId uint8
	CollectorId uint8
    Readings byte
	Stats uint8   //Used by COMMAND to set stats of addressed sensor using SensorId
    Checksum uint8
	Options uint8
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

//Parse packet into LETTProtocol Structure
func GetLETTProtocolPacket(data []byte) (*LETTProtocol ,error) {
	lettproto:=&LETTProtocol{}
	r:=bytes.NewReader(data)
	var packet1 uint32
	binary.Read(r, binary.BigEndian, &packet1)
	lettproto.Header= byte(packet1 >> 24) //top 8 bits
	lettproto.PacketNumber = byte(packet1 >> 16 & 0xff) // 8 bits
	lettproto.SensorId = byte(packet1 >> 8 & 0xff) // 8 bits
	lettproto.CollectorId = byte(packet1 & 0x0fff) //lower 8 bits
	
	var packet2 uint32
	binary.Read(r, binary.BigEndian, &packet2)
	lettproto.Readings= byte(packet2 >> 24) //top 8 bits
	lettproto.Stats = byte(packet2 >> 16 & 0xff) // 8 bits
	lettproto.Checksum = byte(packet2 >> 8 & 0xff) // 8 bits
	lettproto.Options = byte(packet2 & 0x0fff) //lower 8 bits
	
	return lettproto,nil
}

func (lettproto *LETTProtocol) Marshal() []byte{
	buf:=new(bytes.Buffer)
	var packet1 uint32
	packet1 = uint32(lettproto.Header)<<24 | uint32(lettproto.PacketNumber)<<16 | uint32(lettproto.SensorId)<<8 | uint32(lettproto.CollectorId)
	binary.Write(buf, binary.BigEndian, packet1)
	
	var packet2 uint32
	packet2= uint32(lettproto.Readings)<<24 | uint32(lettproto.Stats)<<16 | uint32(lettproto.Checksum)<<8 | uint32(lettproto.Options)
	binary.Write(buf, binary.BigEndian, packet2)
	out:=buf.Bytes()
	
	return out
}


