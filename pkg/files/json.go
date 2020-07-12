package files

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadJSON(path string, to interface{}) {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &to)
	if err != nil {
		panic(err)
	}
}
