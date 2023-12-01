package greeting

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("EMPTY NAME")
	}
	return fmt.Sprintf("Hello %s, welcome to Golang learning", name), nil
}
