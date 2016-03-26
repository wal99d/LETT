package LETT

import(
	"fmt"
	"log"
)

//LETT Packet consists of [Header : Destination Address: Length of Expected Data : Data it self : Checksum to validate the bytes]
type LETT struct{
	Header byte
	DestinationAddress byte
    Length [4]byte
    Data []byte
    Checksum []byte
}

type Packet LETT
type Data []byte

//These are const for Headers type in LETT protocol
const (
	PUSH := 0x01
	PULL := 0x02
	REPUSH := 0x03
)

//This function will handle Packet processing from TCP/IP and return *Data
func (p *Packet) GetData() err, Data{
	//TODO: Extract DATA from Packet struct and return its address
}

//This function will encapsulte Data into LETT/Packet
func (p *Packet) EncapsultePushData(data Data) err {
	p.Header = data[0]
	p.DestinationAddress = data[1]
	p.Length = data[2:6]
	p.Data = data[5:]
	//TODO: generate checksum from bytes and store it in p.Checksum, need to be tested once done
	//p.Checksum = 
	return nil
}

//This function must be called first to sort the Data bytes received from Eth before extracting Data 
func (d Data) BigtoLittleEndian() err{
	//TODO: Handle conversion from Big 2 Little Endian 
}

//This function must be called first to sort the Data bytes received from Eth before extracting Data 
func (d Data) LittletoBigEndian() err{
	//TODO: Handle conversion from Little 2 Big Endian 	
}	