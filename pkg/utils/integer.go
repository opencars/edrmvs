package utils

import (
	"strconv"
)

// Atoi converts string to integer.
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

// Atof converts string to float.
func Atof(lexeme *string) (*float64, error) {
	if lexeme == nil {
		return nil, nil
	}

	result, err := strconv.ParseFloat(*lexeme, 64)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
