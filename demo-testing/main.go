package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
I think there are many different ways for us to generate an id for the given url.
In this demo I just put together a random combination of upper case, lower case, and numbers.
 For this example, I have the id length start at 1 and generate longer length ids if any overlapping occures (see main).
*/
func genID(length int) string {
	cap := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	low := "abcdefghijklmnopqrstuvwxyz"
	num := "0123456789"

	var b strings.Builder
	chars := []rune(cap + low + num)
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
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
			for {
				currID := genID(idLength)
				if _, ok := M[currID]; !ok {
					M[currID] = inputURL
					break
				} else {
					idLength += 1
				}
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
