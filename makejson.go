package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	user := map[string]string{
		"name":    "",
		"address": "",
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your name: ")
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("can't read your name due to %s", err)
	}
	user["name"] = strings.Replace(in, "\n", "", -1)

	fmt.Printf("Enter your address: ")

	in, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("can't read your address due to %s", err)
	}
	user["address"] = strings.Replace(in, "\n", "", -1)

	bArr, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("\n%s\n", bArr)
}
