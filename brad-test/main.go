package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	mapFile, err := os.Open("test.json")
	if err != nil {
		logrus.Errorf("error while opening map file: %v", err)
		os.Exit(1)
	}
	defer mapFile.Close()

	mapJson, err := ioutil.ReadAll(mapFile)
	if err != nil {
		logrus.Errorf("error while reading json file: %v", err)
		os.Exit(1)
	}

	var urlMap map[string]string
	err = json.Unmarshal(mapJson, &urlMap)
	if err != nil {
		logrus.Errorf("error while unmarshalling json: %v", err)
		os.Exit(1)
	}

	// Let's append a value
	//    v ID v     v   URL           v
	urlMap["abc"] = "https://google.com"

	fmt.Printf("%s\n", urlMap)

	// pretend this comes from user input/frontend input
	lookupVal := "Yg0"

	output, exists := urlMap[lookupVal]
	if !exists {
		logrus.WithFields(logrus.Fields{
			"lookupVal": lookupVal,
		}).Warn("couldn't find value in map")
		os.Exit(2)
	}

	logrus.WithFields(logrus.Fields{
		"lookupVal": lookupVal,
		"output":    output,
	}).Infof("Found value!!")
}
