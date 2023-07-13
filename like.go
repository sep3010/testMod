package testMode

import (
	"errors"
	"fmt"
)

func Like(food string) (string, error) {
	if food == "" {
		return "", errors.New("...I don't like anything")
	}

	message := fmt.Sprintf("I like %s.", food)
	return message, nil
}
