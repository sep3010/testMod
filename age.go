package testMode

import (
	"errors"
	"fmt"
)

func MyAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("wrong age")
	}

	message := fmt.Sprintf("I'm, %d years old!", age)
	return message, nil
}
