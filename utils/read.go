package utils

import (
	"io/ioutil"
	"os"
)

func Read(file *os.File) string {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		NewLogger().Error.Println("Failed to read file: ", err)
	}
	return string(b)
}
