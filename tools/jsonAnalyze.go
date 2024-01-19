package tools

import (
	"encoding/json"
	"io/ioutil"
)

func JsonAnalyze() {
	jsonFilePath := `C:\Users\cgq78\Documents\json-out.json`
	outputFilePath := `C:\Users\cgq78\Documents\json-analyze.json`

	// Read JSON file
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON data into an array of maps
	var dataArray []map[string]interface{}
	if err := json.Unmarshal(jsonData, &dataArray); err != nil {
		panic(err)
	}

	// Minimize JSON data
	// minimizedJSON, err := json.Marshal(dataArray)
	// if err != nil {
	// 	panic(err)
	// }

	// Find duplicate entries based on the "KOSTL" key
	duplicates := make(map[string][]map[string]interface{})
	for _, item := range dataArray {
		if kostl, ok := item["KOSTL"].(string); ok {
			duplicates[kostl] = append(duplicates[kostl], item)
		}
	}

	// Filter out non-duplicates
	for key, items := range duplicates {
		if len(items) < 2 {
			delete(duplicates, key)
		}
	}

	// Prepare the output
	outputData, err := json.MarshalIndent(duplicates, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write the output to a file
	err = ioutil.WriteFile(outputFilePath, outputData, 0644)
	if err != nil {
		panic(err)
	}
}
