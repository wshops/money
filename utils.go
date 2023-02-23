package money

import (
	"errors"
	"strconv"
	"strings"
)

func isValidStrAmount(s string) bool {
	dotCount := 0
	// check empty or blank string
	if len(s) == 0 || s == " " || s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		// check empty or blank string
		if s[i] == ' ' {
			return false
		}
		// check if there is more than one dot
		if s[i] == '.' {
			dotCount++
			if dotCount > 1 {
				return false
			}
		}
		// check if there is a non-digit character
		if s[i] != '.' && (s[i] < '0' || s[i] > '9') {
			return false
		}
	}
	return true
}

func parseStrToCentsInt(s string) (int64, error) {
	var decPart int64
	var dolPart int64
	var err error
	// in case of no decimal part
	if !strings.Contains(s, ".") {
		// convert dollar part to cents int
		dolPart, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, errors.New(ErrTypeConversion)
		}
	} else if strings.HasPrefix(s, ".") {
		// append 0 if only one digit for decimal part
		if len(s[1:]) == 1 {
			s += "0"
		}
		// convert decimal part to cents int
		decPart, err = strconv.ParseInt(s[1:], 10, 64)
		if err != nil {
			return 0, errors.New(ErrTypeConversion)
		}
	} else {
		// in case of dollar and decimal part
		splitAmount := strings.Split(s, ".")
		// append 0 if only one digit for decimal part
		if len(splitAmount[1]) == 1 {
			splitAmount[1] += "0"
		}
		// convert decimal part to cents int
		decPart, err = strconv.ParseInt(splitAmount[1], 10, 64)
		if err != nil {
			return 0, errors.New(ErrTypeConversion)
		}
		// convert dollar part to cents int
		dolPart, err = strconv.ParseInt(splitAmount[0], 10, 64)
		if err != nil {
			return 0, errors.New(ErrTypeConversion)
		}
	}
	// calculate amount in cents
	return dolPart*100 + decPart, nil
}
