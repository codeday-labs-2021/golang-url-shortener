package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	// "strings"
	"net/url"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func genID(length int) string {

	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)
	return id
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

/*
in our case, the dictionary will probably be turned into a database or something.
*/
func main() {
	rand.Seed(time.Now().UnixNano())
	M := make(map[string]string)
	idLength := 1
	for {
		fmt.Println("")
		fmt.Println("Type (c/C) to check the stored data for this run")
		fmt.Println("Type (s/S) to shorten/store a url")
		fmt.Println("Type (l/L) to loop up a url using a key/id")
		fmt.Print("Type anything else to quit:  ")

		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		input := sc.Text()
		fmt.Println("")
		if input == "c" || input == "C" {
			for id, url := range M {
				fmt.Println("ID:", id, "=> URL:", url)
			}
		} else if input == "s" || input == "S" {
			fmt.Println("Please paste your url: ")
			sc.Scan()
			inputURL := sc.Text()
			if isUrl(inputURL) {
				for {
					currID := genID(idLength)
					if _, ok := M[currID]; !ok {
						M[currID] = inputURL
						break
					} else {
						idLength += 1
					}
				}
			} else {
				fmt.Println("invalid url")
			}
		} else if input == "l" || input == "L" {
			fmt.Println("Please paste an ID: ")
			sc.Scan()
			inputID := sc.Text()
			if val, ok := M[inputID]; ok {
				fmt.Println("Your Url is: " + val)
			} else {
				fmt.Println("The ID you provided is invalid")
			}
		} else {
			break
		}
	}
}
