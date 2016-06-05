package LETT

import(
	"testing"
	"bytes"
	"encoding/gob"
	"net"
)

func TestNewLETTPacket(t *testing.T){

	
	lettproto:= &LETTProtocol{
		Header: COMMAND,
		PacketNumber: 1,
		SensorId: 1,
		Stats: OFF,
	}
	
	var conn *net.Conn = nil
	
	NewLETTPacket(conn, lettproto)
}

func TestGetLETTProtocolPacket(t *testing.T){
	
	lettproto:= &LETTProtocol{
		Header: COMMAND,
		PacketNumber: 1,
		SensorId: 1,
		Stats: ON,
	}
	
	buf:= &bytes.Buffer{}
	//buf:= bytes.NewBuffer([]byte(lettproto))
	err:= gob.NewEncoder(buf).Encode(lettproto)
	if err != nil {
		panic(err)
	}
	var conn *net.Conn
	GetLETTProtocolPacket(conn)
	
}
