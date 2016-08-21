package signed

import "testing"

func TestKey(t *testing.T) {
	SetSecretKey("1234")
	if len(secretKey) != 32 {
		t.Error("SetSecretKey function error")
	}
}

func TestSigned(t *testing.T) {
	text := "1234567891234567"
	e := Encode(text)
	d := Decode(e)
	if text != d {
		t.Error("signed error")
	}
}
