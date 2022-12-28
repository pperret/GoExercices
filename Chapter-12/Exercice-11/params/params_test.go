package params

import (
	"net/url"
	"testing"
)

func TestPack(t *testing.T) {

	data := struct {
		MyString     string `http:"ms"`
		MyInteger    int    `http:"mi"`
		ListElements []int  `http:"ei"`
	}{"str", 3, []int{5, 8, 10}}

	u, _ := url.Parse("http://www.gopl.io")
	err := Pack(u, &data)
	if err != nil {
		t.Errorf("Unable to pack parameters: %v", err)
	}
	expected := "http://www.gopl.io?ei=5&ei=8&ei=10&mi=3&ms=str"
	if u.String() != expected {
		t.Errorf("Bad encoding: expected=%v, got=%v", expected, u.String())
	}
	t.Logf("url=%v", u.String())
}
