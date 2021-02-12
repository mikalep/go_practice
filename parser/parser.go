/*
*  JSON parser that parses a file based on search query. Subject to change.
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Personnel struct {
	Salaries []Salaries `json:"salaries"`
}

type Salaries struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

func main() {

	fmt.Println("Which file would you like to parse?")

	scanResult, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You have to give a file to parse.")
		os.Exit(-1)
	}

	result, err := ioutil.ReadAll(scanResult)
	if err != nil {
		log.Fatal(err)
	}

	var personnel Personnel

	err = json.Unmarshal(result, &personnel)
	if err != nil {
		log.Fatal(err)
	}

	json, err := json.MarshalIndent(personnel, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
