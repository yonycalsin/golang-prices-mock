package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"path"

	"github.com/labstack/echo"
)

// PriceLists is for payload
type PriceLists struct {
	Data []PriceList `json:"data"`
}

// PriceList Type
type PriceList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Variants ...
type Variants struct {
	Data []Variant `json:"data"`
}

// Variant ...
type Variant struct {
	ID         int                `json:"id"`
	SKU        string             `json:"sku"`
	Attributes []VariantAttribute `json:"attributes"`
}

// VariantAttribute ...
type VariantAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Prices ...
type Prices struct {
	Data []Price `json:"data"`
}

// Price ...
type Price struct {
	ID          int    `json:"id"`
	Price       string `json:"price"`
	OldPrice    string `json:"oldPrice"`
	RibbonLabel string `json:"ribbonLabel"`
	ProductID   int    `json:"ProductId"`
}

var currentFolder, err = os.Getwd()

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
	mockFilename := "/data/price-lists.json"

	filename := path.Join(currentFolder, mockFilename)

	bytes := GetMockInBytes(filename)

	var payload PriceLists

	json.Unmarshal(bytes, &payload)

	return c.JSON(http.StatusOK, payload)
}

func variantsHandler(c echo.Context) error {
	mockFilename := "/data/variants.json"

	filename := path.Join(currentFolder, mockFilename)

	bytes := GetMockInBytes(filename)

	var payload Variants

	json.Unmarshal(bytes, &payload)

	return c.JSON(http.StatusOK, payload)
}

func pricesHandler(c echo.Context) error {
	mockFilename := "/data/prices.json"

	filename := path.Join(currentFolder, mockFilename)

	bytes := GetMockInBytes(filename)

	var payload Prices

	json.Unmarshal(bytes, &payload)

	return c.JSON(http.StatusOK, payload)
}

func main() {

	server := echo.New()

	server.GET("/", mainHandler)

	server.GET("/api/price-lists", priceListsHandler)

	server.GET("/api/variants", variantsHandler)

	server.GET("/api/prices", pricesHandler)

	server.Logger.Fatal(server.Start(":4444"))
}
