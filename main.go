package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
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

/* Handlers */

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func priceListsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hey from PriceLists Endpoint")
}

func main() {
	fmt.Println("Hello World")

	server := echo.New()

	server.GET("/", mainHandler)

	server.GET("/api/price-lists", priceListsHandler)

	server.Logger.Fatal(server.Start(":4444"))
}
