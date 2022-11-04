package prototype

import (
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	s1 := Shape{
		x:     1,
		y:     2,
		color: "red",
	}
	s2 := s1.Clone()
	if !reflect.DeepEqual(s1, s2) {
		t.Fatal("Error")
	}
}
