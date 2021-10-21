package simplefactory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	output := api.Say("Zhening")
	if output != "Hi, Zhening" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	output := api.Say("Zhening")
	if output != "Hello, Zhening" {
		t.Fatal("Type2 test fail")
	}
}

func TestTypeNil(t *testing.T) {
	api := NewAPI(0)
	if api != nil {
		t.Fatal("TypeNil test fail")
	}
}
