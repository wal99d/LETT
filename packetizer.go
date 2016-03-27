package LETT

import(
	//"fmt"
	"errors"
)

//LETT Packet consists of [Header : Destination Address: Length of Expected Data : Data it self : Checksum to validate the bytes]
type LETTPacket struct{
	Header byte
	DestinationAddress byte
    Length []byte
    Data []byte
    Checksum []byte
}

type Data []byte

//These are const for Headers type in LETT protocol , maximum is 2^8 = 256 values
const (
	PUSH = 0x01
	PULL = 0x02
	REPUSH = 0x03
	REPULL =0x04
)

//LETT Defined Errors
var(
	NoData = errors.New("There is no available Data on the stream yet")
)

//This function will handle Packet processing from TCP/IP and return *Data
func (p *LETTPacket) GetData() (error, Data) {
	if p.Data == nil {
		return NoData , nil
	}else{
		return nil, p.Data		
	}

}

//This function will encapsulte Data into LETT/Packet
func (p *LETTPacket) EncapsulteData(data Data) error {
	if data !=nil{
		p.Header = data[0]
		p.DestinationAddress = data[1]
		p.Length = data[2:6]
		p.Data = data[5:]
		checksum:=make([]byte, len(data))
		for k,v :=range data{
			checksum[k]+= v
			checksum[k]&=0xff
		}
		p.Checksum = checksum
		return nil
	}else{
		return NoData
	}
}

//This function must be called first to sort the Data bytes received from Ethernet before extracting Data 
func (d Data) BigtoLittleEndian() error{
	//TODO: Handle conversion from Big 2 Little Endian 
	return nil
}

//This function must be called first to sort the Data bytes received from Ethernet before extracting Data 
func (d Data) LittletoBigEndian() error{
	//TODO: Handle conversion from Little 2 Big Endian 
	return nil	
}	