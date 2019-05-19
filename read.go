package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

var persons []Person

const maxSize = 20

func main() {

	var fileName string
	fmt.Printf("Enter File Name : ")
	if _, err := fmt.Scan(&fileName); err != nil {
		log.Fatalf("Can't read fileName, reason: %v\n", err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Can't open file named '%s'\n", fileName)
	}

	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')

		slice := strings.Split(strings.Replace(str, "\n", "", -1), " ")

		var firstName, lastName string
		if len(slice) > 1 {
			firstName = slice[0]
			lastName = slice[1]
		} else if len(slice) == 1 {
			firstName = slice[0]
			lastName = ""
		}

		if len(firstName) > maxSize {
			firstName = firstName[0:maxSize]
		}

		if len(lastName) > maxSize {
			lastName = lastName[0:maxSize]
		}

		persons = append(persons, Person{
			firstName: firstName,
			lastName:  lastName,
		})

		if err != nil {
			break
		}
	}

	if err = f.Close(); err != nil {
		log.Fatalf("Can't close file '%s', reason: %v\n", fileName, err)
	}

	if len(persons) > 0 {
		for i, j := range persons {
			fmt.Printf("%d) FirstName: '%s', LastName: '%s'\n", i+1, j.firstName, j.lastName)
		}
	}
}
