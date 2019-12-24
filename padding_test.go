package padding

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var tString = []byte("test")

func TestPKCS5_Pad(t *testing.T) {
	pkcs5 := NewPKCS5()
	result := pkcs5.Pad(tString, 8)
	s := base64.StdEncoding.EncodeToString(result)
	if s != "dGVzdAQEBAQ=" {
		t.Errorf("expect: `dGVzdAQEBAQ=`, got: %s", s)
	}
}

func TestPKCS5_UnPad(t *testing.T) {
	pkcs5 := NewPKCS5()
	b, err := base64.StdEncoding.DecodeString("dGVzdAQEBAQ=")
	if err != nil {
		t.Errorf("%v", err)
	}
	r, err := pkcs5.UnPad(b)
	if err != nil {
		t.Errorf("%v", err)
	}
	if fmt.Sprintf("%s", r) != "test" {
		t.Errorf("expect: test, got: %s", fmt.Sprintf("%s", r))
	}
}

func TestZero_Pad(t *testing.T) {
	z := NewZero()
	r := z.Pad(tString, 8)
	s := base64.StdEncoding.EncodeToString(r)
	if s != "dGVzdAAAAAA=" {
		t.Errorf("expect: `dGVzdAAAAAA=` , got: %s", s)
	}
}
