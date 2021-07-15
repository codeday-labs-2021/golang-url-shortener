package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func genID(length int) string {

	id, err := gonanoid.New(length)
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

type urldatabase struct {
	UrlID   string `json:"urlID"`
	LongURL string `json:"longURL"`
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Type 'r' to run server or 'a' to add urls to the database:  ")
	sc.Scan()
	in := sc.Text()

	if in == "r" || in == "R" {
		var dataArray []urldatabase

		jsonFile, err := os.Open("urlmap.json")
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		if err := json.Unmarshal(byteValue, &dataArray); err != nil {
			log.Println(err)
		}

		jsonFile.Close()
		r := gin.Default()

		r.GET("/home", func(c *gin.Context) {
			c.JSON(200, gin.H{"hello": "world"})
		})

		for i := range dataArray {
			r.GET(dataArray[i].UrlID, func(c *gin.Context) {
				c.Redirect(http.StatusMovedPermanently, dataArray[i].LongURL)
			})
		}
		for _, v := range dataArray {
			fmt.Println("localhost:8080/" + v.UrlID + "          " + v.LongURL)
		}
		r.Run(":8080")

	} else {
		var dataArray []urldatabase

		jsonFile, err := os.Open("urlmap.json")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened url.json")
		byteValue, _ := ioutil.ReadAll(jsonFile)
		if err := json.Unmarshal(byteValue, &dataArray); err != nil {
			log.Println(err)
		}

		M := make(map[string]string)
		idLength := 3
		for {

			fmt.Println("")
			fmt.Println("Type (c/C) to check the stored data for this run")
			fmt.Println("Type (s/S) to shorten/store a url")
			fmt.Println("Type (l/L) to look up a url using a key/id")
			fmt.Print("Type anything else to quit:  ")

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
							dataArray = append(dataArray, urldatabase{UrlID: currID, LongURL: inputURL})
							for _, v := range dataArray {
								fmt.Println(v)
							}
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

		result, err := json.Marshal(dataArray)
		if err != nil {
			log.Println(err)
		}
		ioutil.WriteFile("urlmap.json", result, 0644)
		jsonFile.Close()
	}

}
