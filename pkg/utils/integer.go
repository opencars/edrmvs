package utils

import (
	"strconv"
)

func Atoi(lexeme *string) (*int, error) {
	if lexeme == nil {
		return nil, nil
	}

	result, err := strconv.Atoi(*lexeme)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
