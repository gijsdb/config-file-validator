package validator

import "encoding/xml"

func NewXMLValidator() InteractorI {
	return &XMLValidator{}
}

type XMLValidator struct {
}

func (xv XMLValidator) Validate(b []byte) (bool, error) {
	var output interface{}
	err := xml.Unmarshal(b, &output)
	if err != nil {
		return false, err
	}
	return true, nil
}
