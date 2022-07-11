package util

import (
	"encoding/json"
	"io/ioutil"
)

func WriteJsonFile(filePath string, data any) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	ioutil.WriteFile(filePath, file, 0644)
	return nil
}
