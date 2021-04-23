package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {

	if IsEmpty(fileName) {
		return "", errors.New("FileName cannot be null.")
	}

	bytes, err := ioutil.ReadFile(fileName)
	CheckErrors(err)
	return string(bytes), err
}
