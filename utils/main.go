package utils

import (
	"encoding/json"
	"os"
)

type Data struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	IsMarked bool   `json:"isMarked"`
	Priority int    `json:"priority"`
}

func Read(filename string) []Data {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var content []Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&content)
	if err != nil {
		panic(err)
	}
	return content
}
