package ptrutil

import (
	"reflect"
	"testing"
)

type TestStruct struct{}

func TestPtr(t *testing.T) {
	tests := map[string]interface{}{
		"string": "test",
		"int":    123,
		"bool":   true,
		"rune":   'B',
		"struct": TestStruct{},
	}

	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Ptr(val); !reflect.DeepEqual(got, &val) {
				t.Errorf("Ptr() = %v, want %v", got, &val)
			}
		})
	}
}

func TestFromPtr_NilString(t *testing.T) {
	var val *string

	res := FromPtr(val)

	if res != "" {
		t.Error("should be empty string")
	}
}

func TestFromPtr_NilInt(t *testing.T) {
	var val *int

	res := FromPtr(val)

	if res != 0 {
		t.Error("should be 0")
	}
}

func TestFromPtr_String(t *testing.T) {
	val := "test"
	res := FromPtr(&val)

	if res != val {
		t.Error("should equal")
	}
}
