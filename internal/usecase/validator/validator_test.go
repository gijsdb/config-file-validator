package validator

import (
	"testing"
)

var testData = []struct {
	name           string
	testInput      []byte
	expectedResult bool
	validator      InteractorI
}{
	{"validJson", []byte(`{"test": "test"}`), true, JSONValidator{}},
	{"invalidJson", []byte(`{test": "test"}`), false, JSONValidator{}},
	{"validYaml", []byte("a: 1\nb: 2"), true, YAMLValidator{}},
	{"invalidYaml", []byte("a: b\nc: d:::::::::::::::"), false, YAMLValidator{}},
	{"validXml", []byte("<test>\n</test>"), true, XMLValidator{}},
	{"invalidXml", []byte("<xml\n"), false, XMLValidator{}},
	{"invalidToml", []byte("name = 123__456"), false, TOMLValidator{}},
	{"validToml", []byte("name = 123"), true, TOMLValidator{}},
}

func Test_ValidationInput(t *testing.T) {
	for _, d := range testData {
		valid, err := d.validator.Validate(d.testInput)
		if valid != d.expectedResult {
			t.Errorf("incorrect result: expected %v, got %v", d.expectedResult, valid)
		}

		if valid && err != nil {
			t.Error("incorrect result: err was not nil", err)
		}

		if !valid && err == nil {
			t.Error("incorrect result: function returned a nil error")
		}
	}
}
