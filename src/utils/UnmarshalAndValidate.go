package utils

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

func UnmarshalAndValidate(data []byte, v any) error {

	err := json.Unmarshal(data, v)

	if err != nil {
		return err
	}

	validate := validator.New()

	return validate.Struct(v)
}
