package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"path"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	PriceListID int    `json:"PriceListId"`
}

var currentFolder, _ = os.Getwd()

// GetMockInBytes will return json data file for mock
func GetMockInBytes(filename string) []byte {
	mockFilename := path.Join(currentFolder, filename)

	jsonFile, err := os.Open(mockFilename)

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
	bytes := GetMockInBytes("/data/price-lists.json")

	var payload PriceLists

	json.Unmarshal(bytes, &payload)

	return c.JSON(http.StatusOK, payload)
}

func variantsHandler(c echo.Context) error {
	bytes := GetMockInBytes("/data/variants.json")

	var payload Variants

	json.Unmarshal(bytes, &payload)

	return c.JSON(http.StatusOK, payload)
}

func pricesHandler(c echo.Context) error {
	bytes := GetMockInBytes("/data/prices.json")

	var payload Prices

	json.Unmarshal(bytes, &payload)

	return c.JSON(http.StatusOK, payload)
}

func main() {
	server := echo.New()

	// Loggin
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	loggerMiddlware := NewLoggerMiddleware(logger)

	server.Logger = loggerMiddlware

	server.Use(loggerMiddlware.Hook())

	// Middlewares ...
	server.Use(middleware.Recover())

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Routes
	server.GET("/", mainHandler)

	server.GET("/api/price-lists", priceListsHandler)

	server.GET("/api/variants", variantsHandler)

	server.GET("/api/prices", pricesHandler)

	server.Logger.Fatal(server.Start(":4444"))
}
