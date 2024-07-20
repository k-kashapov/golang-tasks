//go:build !solution

package mycheck

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type errorSlice []error

var (
	numErr  = errors.New("found numbers")
	longErr = errors.New("line is too long")
	spcErr  = errors.New("no two spaces")
)

func (es errorSlice) Error() string {
	var sb strings.Builder
	for _, err := range es {
		if err != nil {
			if sb.Len() > 0 {
				sb.WriteString(";")
			}

			sb.WriteString(fmt.Sprintf("%v", err))
		}
	}

	return sb.String()
}

func MyCheck(input string) error {
	es := make(errorSlice, 3)
	num_err := 0

	num_space := 0

	for _, ch := range input {
		if unicode.IsDigit(ch) {
			num_err++
			es[0] = numErr
		}

		if ch == ' ' {
			num_space++
		}
	}

	if len(input) > 20 {
		num_err++
		es[1] = longErr
	}

	if num_space != 2 {
		num_err++
		es[2] = spcErr
	}

	if num_err != 0 {
		return es
	}

	return nil
}
