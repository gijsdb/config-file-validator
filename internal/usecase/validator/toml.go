package validator

import (
	"errors"
	"fmt"

	"github.com/pelletier/go-toml/v2"
)

func NewTOMLValidator() InteractorI {
	return &TOMLValidator{}
}

type TOMLValidator struct {
}

func (tv TOMLValidator) Validate(b []byte) (bool, error) {
	var output interface{}
	err := toml.Unmarshal(b, &output)
	var derr *toml.DecodeError
	if errors.As(err, &derr) {
		row, col := derr.Position()
		return false, fmt.Errorf("error at line %v column %v: %v", row, col, err)
	}
	return true, nil
}
