package LETT

import(
	"testing"
	"bytes"
)

func TestGetHeader(t *testing.T){
	p:=LETTPacket{}
	buf := []byte{0x01,0x01,0xff,0x7f,0x00,0x00,0x11,0x11,0x11,0x11}
	p.Buffer = bytes.NewReader(buf)

	if err, _:= p.GetHeader(); err!=nil{
		t.Fatalf("%s\n", err)
	}
}

func TestGetDestinationAddress(t *testing.T){
	p:=LETTPacket{}
	buf := []byte{0x01,0x01,0xff,0x7f,0x00,0x00,0x11,0x11,0x11,0x11}
	p.Buffer = bytes.NewReader(buf)

	if err, _:= p.GetDestinationAddress(); err!=nil{
		t.Fatalf("%s\n", err)
	}
}

func TestGetData(t *testing.T){
	p:=LETTPacket{}
	buf := []byte{0x01,0x01,0xff,0x7f,0x00,0x00,0x11,0x11,0x11,0x11}
	p.Buffer = bytes.NewReader(buf)

	if err, _:= p.GetData(); err!=nil{
		t.Fatalf("%s\n", err)
	}
}

func TestGetChecksum(t *testing.T){
	p:=LETTPacket{}
	buf := []byte{0x01,0x01,0xff,0x7f,0x00,0x00,0x11,0x11,0x11,0x11}
	p.Buffer = bytes.NewReader(buf)

	if err, _:= p.GetChecksum(); err!=nil{
		t.Fatalf("%s\n", err)
	}
}

func TestGetPacket(t *testing.T){
	p:=LETTPacket{}
	buf := []byte{0x01,0x01,0xff,0x7f,0x00,0x00,0x11,0x11,0x11,0x11}
	p.Buffer = bytes.NewReader(buf)

	if err, _:= p.GetPacket(); err!=nil{
		t.Fatalf("%s\n", err)
	}
}

func TestNewLETTPacket(t *testing.T){
	buf := []byte{0x01,0x01,0xff,0x7f,0x00,0x00,0x11,0x11,0x11,0x11}
	buffer :=bytes.NewReader(buf)
	NewLETTPacket(buffer)
}