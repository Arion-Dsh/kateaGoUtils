package signed

import "testing"

func TestKey(t *testing.T) {
	s := Signed{}
	s.SetSecretKey("12345")
	if len(s.secretKey) != 32 {
		t.Error("SetSecretKey function error")
	}
}

func TestSigned(t *testing.T) {
	s := Signed{}
	s.SetSecretKey("12345")
	text := "1234567"
	e := s.AESEncode(text)
	d := s.AESDecode(e)
	if text != d {
		t.Error("signed error")
	}
}
