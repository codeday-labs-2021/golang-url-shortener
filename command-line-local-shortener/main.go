package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sirupsen/logrus"
)

var dataArray map[string]string

func genID(length int) string {

	id, err := gonanoid.New(length)
	if err != nil {
		logrus.Errorf("error while creating new id: %v", err)
		os.Exit(1)
	}
	return id
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		logrus.Errorf("error parsing url: %v", err)
		os.Exit(1)
	}
	return u.Scheme != "" && u.Host != ""
}

func dataProccess() {
	sc := bufio.NewScanner(os.Stdin)
	idLength := 3
	for {

		sc.Scan()
		input := sc.Text()
		fmt.Println("")

		if input == "s" || input == "S" {
			fmt.Println("Please paste your url: ")
			sc.Scan()
			inputURL := sc.Text()
			if isUrl(inputURL) {
				for {
					currID := genID(idLength)
					if _, exists := dataArray[currID]; !exists {
						dataArray[currID] = inputURL
						result, err := json.Marshal(dataArray)
						if err != nil {
							logrus.Errorf("error while marshalling json: %v", err)
							os.Exit(1)
						}
						ioutil.WriteFile("urlmap.json", result, 0644)

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
		} else {
			break
		}
	}
}

func getReq(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)
	url, exists := dataArray[id]
	if exists {
		c.Redirect(http.StatusMovedPermanently, url)
	}
}

func main() {

	jsonFile, err := os.Open("urlmap.json")
	if err != nil {
		logrus.Errorf("error while opening map file: %v", err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened url.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &dataArray); err != nil {
		logrus.Errorf("error while unmarshalling json: %v", err)
		os.Exit(1)
	}

	dataProccess()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/:id", getReq)
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")

}
