package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetConfig() Configuration {
	absPath, _ := filepath.Abs("./config.json")

	// Read JSON-file
	jsonFile, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		return Configuration{}
	}
	defer jsonFile.Close()

	// Read json content
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Parse JSON
	var configuration Configuration
	json.Unmarshal(byteValue, &configuration)

	return configuration
}
