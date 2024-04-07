package simplefactory

import "testing"

// test Type1 say Hi
func TestType1(t *testing.T) {
	api := NewAPI(1)

	if s := api.Say("Tom"); s != "Hi, Tom" {
		t.Fatal("Type1 1.test_AAA fail")
	}
}

// test Type2 say Hello
func TestType2(t *testing.T) {
	api := NewAPI(2)

	if s := api.Say("Mike"); s != "Hello, Mike" {
		t.Fatal("Type2 1.test_AAA fail")
	}
}
