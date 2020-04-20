package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func main() {

	schemaLoader := gojsonschema.NewReferenceLoader("file://./schema.json")
	documentLoader1 := gojsonschema.NewReferenceLoader("file://./example1.json")
	documentLoader2 := gojsonschema.NewReferenceLoader("file://./example2.json")

	result1, err := gojsonschema.Validate(schemaLoader, documentLoader1)
	result2, err := gojsonschema.Validate(schemaLoader, documentLoader2)
	if err != nil {
		panic(err.Error())
	}

	if result1.Valid() {
		fmt.Printf("The LoRaWAN FUOTA manifest example 1 is valid\n")
	} else {
		fmt.Printf("The LoRaWAN FUOTA manifest example 2 is not valid. see errors :\n")
		for _, err := range result1.Errors() {
			// Err implements the ResultError interface
			fmt.Printf("- %s\n", err)
		}
	}
	if result2.Valid() {
		fmt.Printf("The LoRaWAN FUOTA manifest example 2 is valid\n")
	} else {
		fmt.Printf("The LoRaWAN FUOTA manifest example 2 is not valid. see errors :\n")
		for _, err := range result2.Errors() {
			// Err implements the ResultError interface
			fmt.Printf("- %s\n", err)
		}
	}
}
