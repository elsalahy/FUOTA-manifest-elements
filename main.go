package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
		example1JsonFile, err := os.Open("example1.json")
		if err != nil {
			fmt.Println(err)
		}
		example1Binary, err := os.Create("example1.bin")
		if err != nil {
			fmt.Println(err)
		}
		jsonData, _ := ioutil.ReadAll(example1JsonFile)
		example1Manifest := manifest{}
		json.Unmarshal(jsonData, &example1Manifest)
		writeBinary(&example1Manifest, example1Binary)
		example1Binary.Close()
		fmt.Printf("Successfully generated binary file for example1, see example1.bin\n")
	} else {
		fmt.Printf("The LoRaWAN FUOTA manifest example 1 is not valid. see errors :\n")
		for _, err := range result1.Errors() {
			// Err implements the ResultError interface
			fmt.Printf("- %s\n", err)
		}
	}
	if result2.Valid() {
		fmt.Printf("The LoRaWAN FUOTA manifest example 2 is valid\n")
		example2JsonFile, err := os.Open("example2.json")
		if err != nil {
			fmt.Println(err)
		}
		example2Binary, err := os.Create("example2.bin")
		if err != nil {
			fmt.Println(err)
		}
		jsonData, _ := ioutil.ReadAll(example2JsonFile)
		example2Manifest := manifest{}
		json.Unmarshal(jsonData, &example2Manifest)
		writeBinary(&example2Manifest, example2Binary)
		example2Binary.Close()
		fmt.Printf("Successfully generated binary file for example2, see example2.bin\n")
	} else {
		fmt.Printf("The LoRaWAN FUOTA manifest example 2 is not valid. see errors :\n")
		for _, err := range result2.Errors() {
			// Err implements the ResultError interface
			fmt.Printf("- %s\n", err)
		}
	}
}
