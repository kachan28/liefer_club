package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	protokolsFolder    = "LC-Kasse-Programmierprotokolle"
	protokolMenuFolder = "JSON-Format"
)

var (
	protokolPath = filepath.Join(protokolsFolder, protokolMenuFolder)
)

func main() {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(filepath.Join(currPath, protokolPath)); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Join(currPath, protokolPath), 0755)
		if err != nil {
			panic(err)
		}
	}
	jsonResult, err := json.Marshal(struct {
		name    string `json:"name"`
		surname string `json:"surname"`
	}{
		name:    "sad",
		surname: "asd",
	})
	if err != nil {
		panic(err)
	}
	protocolFileCreationTime := time.Now()
	protocolFileName := fmt.Sprintf("%s_%s", "create_attempt", protocolFileCreationTime.Format("2006.01.02_15-04-05"))
	err = os.WriteFile(filepath.Join(currPath, protokolPath, protocolFileName)+".json", jsonResult, 0644)
	// err = os.WriteFile("temp.json", jsonResult, 0644)
	if err != nil {
		panic(err)
	}
}
