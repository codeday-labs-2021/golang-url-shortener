package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	// "time"
	"net/http"

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

/*
in our case, the dictionary will probably be turned into a database or something.
*/
func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Type 'r' to run server or 'a' to add urls to the database:  ")
	sc.Scan()
	in := sc.Text()

	if in == "r" || in == "R" {

		r := gin.Default()

		r.GET("/test", func(c *gin.Context) {
			c.Request.URL.Path = "/test2"
			r.HandleContext(c)
		})
		r.GET("/test2", func(c *gin.Context) {
			c.JSON(200, gin.H{"hello": "world"})
		})

		r.GET("/test3", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
		})

		r.Run(":8080")

	} else {

		M := make(map[string]string)
		idLength := 3
		for {
			fmt.Println("")
			fmt.Println("Type (c/C) to check the stored data for this run")
			fmt.Println("Type (s/S) to shorten/store a url")
			fmt.Println("Type (l/L) to loop up a url using a key/id")
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
							// r.GET(currID, func(c *gin.Context) {
							// 	c.Redirect(http.StatusMovedPermanently, inputURL)
							// })

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

}
