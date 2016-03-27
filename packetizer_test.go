package LETT

import(
	"testing"
)

func TestEncapsulteDataGetData(t *testing.T){
	p:=LETTPacket{}
	data := []byte{0x01,0x01,0xff,0x7f,0x00,0x00,0x11,0x11,0x11}
	if err:= p.EncapsulteData(data); err!=nil {
		t.Fatalf("%s\n", err)
	}
	if err, d:= p.GetData(); err!=nil {
		t.Fatalf("%s\n", err)
	}
}