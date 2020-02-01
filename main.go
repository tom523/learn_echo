package learn_echo

import (
	"errors"
	"fmt"
)

// Hi returns a friendly greeting
func Hi(name string, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "cn":
		return fmt.Sprintf("NiHao, %s!", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
