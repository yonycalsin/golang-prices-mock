package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetMockInBytes will return json data file for mock
func GetMockInBytes(filename string) []byte {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully opened", filename)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func main() {
	fmt.Println("Hello World")
}
