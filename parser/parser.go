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
	"strings"
)

type Personnel struct {
	Salaries []Salaries `json:"salaries"`
}

type Salaries struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

func handleSearchTerms(p Personnel) {

	fmt.Println("Would you like to search with n(ame) or with a s(alary)?")

	var choice string

	fmt.Scanln(&choice)

	if choice == "n" || choice == "name" {

		name := searchByName()

		formattedName := strings.ToLower(name)

		for _, v := range p.Salaries {

			if formattedName == strings.ToLower(v.Name) {
				fmt.Printf("Found %s, with a salary of: %d\n", v.Name, v.Salary)
			}
		}

	} else if choice == "s" || choice == "salary" {

		salary := searchBySalary()

		for _, v := range p.Salaries {
			if salary == v.Salary {
				fmt.Printf("Found %s, with a salary of: %d\n", v.Name, v.Salary)
			}
		}
	} else {
		fmt.Println("Not a valid search parameter, would you like to try again?")
	}
}

func searchByName() string {
	var name string

	fmt.Println("Filter for name(s): ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		log.Fatal(err)
	}

	return name
}

func searchBySalary() int {
	var salary int

	fmt.Println("Please give a minimal salary: ")
	_, err := fmt.Scanln(&salary)

	if err != nil {
		log.Fatal(err)
	}

	return salary
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You have to give a file to parse.")
		os.Exit(-1)
	}

	scanResult, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
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

	handleSearchTerms(personnel)
}
