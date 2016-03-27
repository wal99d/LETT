package LETT

import(
	//"fmt"
	"errors"
	"encoding/binary"
	//"bytes"
	"io/ioutil"
	"io"
)

//LETT Packet consists of [Header : Destination Address: Data itself of 2 Bytes in Size : Checksum to validate the bytes]
type LETTPacket struct{
	Buffer io.Reader
}

type LETTProtocol struct{
	Header uint8
	DestinationAddress uint8
    Data uint16
    Checksum uint32
}

//These are const for Headers type in LETT protocol , maximum is 2^8 = 256 values
const (
	READ = 0x01
	COMMAND = 0x02
)

//LETT Defined Errors
var(
	EmptyBuffer = errors.New("Opps Buffer is Empty!!")
)

//This function will check for Buffer and return error if it's nil and []byte
func (p *LETTPacket) CheckBuffer() (error, []byte) {
	if packet, err:=ioutil.ReadAll(p.Buffer); err!=nil && err == io.EOF{
		return EmptyBuffer , nil
	}else{
		return nil, packet		
	}
}

//This function will handle Packet processing from TCP/IP and return Buffer/[]byte
func (p *LETTPacket) GetPacket() (error, []byte ) {
	if err, packet :=p.CheckBuffer(); err ==nil{
		return nil , packet
	}else {
		return err , nil
	}
}

//This function will return the header from received buffer as uint8
func (p *LETTPacket) GetHeader() (error , uint8){
	if err, packet :=p.CheckBuffer(); err ==nil{
		return nil , packet[0]
	}else {
		return err , 0
	}
}

//This function will return the DestinationAddress from received buffer as uint8
func (p *LETTPacket) GetDestinationAddress() (error , uint8){
	if err, packet :=p.CheckBuffer(); err ==nil{
		return nil , packet[1]
	}else {
		return err , 0
	}
}

//This function is used to get the Data form Buffer
func (p *LETTPacket) GetData() (error , uint16){
	if err, packet :=p.CheckBuffer(); err ==nil{
		return nil , binary.BigEndian.Uint16(packet[2:4])
	}else {
		return err , 0
	}
}

//This function is used to get the checksum form Buffer
func (p *LETTPacket) GetChecksum() (error , uint32){
	if err, packet :=p.CheckBuffer(); err ==nil{
		return nil , binary.BigEndian.Uint32(packet[4:])
	}else {
		return err , 0
	}
}
