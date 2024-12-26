package utils

import "errors"

func PadLeft(str string, length int) (string, error) {
	if len(str) > length {
		return "", errors.New("String is too long to pad!")
	}

	for len(str) < length {
		str = " " + str
	}

	return str, nil
}

func PadRight(str string, length int) (string, error) {
	if len(str) > length {
		return "", errors.New("String is too long to pad!")
	}

	for len(str) < length {
		str = str + " "
	}

	return str, nil
}
