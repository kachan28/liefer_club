package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const mariaDbConfFile = "dbconnection.json"

type connection struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (c *connection) getConf() {
	jsonFile, err := os.Open(mariaDbConfFile)
	defer jsonFile.Close()
	if err != nil {
		log.Fatal("can't open mariaDB config", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, c)
	if err != nil {
		log.Fatal("can't unmarshal json config", err)
	}
}
