package mjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ManifestInfo(path string) string {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	r := ""
	for k, v := range result {
		if k == "name" {
			r += fmt.Sprintf("%s\n", v)
		}
	}
	for k, v := range result {
		if k == "version" {
			r += fmt.Sprintf("%s\n", v)
		}
	}
	for k, v := range result {
		if k == "update_url" {
			r += fmt.Sprintf("%s\n", v)
		}
	}
	return r
}
