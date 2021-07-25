package main

import (
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

func dataProccess(inputURL string) string {
	idLength := 3
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

				return currID
			} else {
				idLength += 1
			}
		}
	}
	return ""

}

func getReq(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)
	url, exists := dataArray[id]
	if exists {
		c.Redirect(http.StatusMovedPermanently, url)
	}
}

func postReq(c *gin.Context) {
	longURL := c.PostForm("longURL")

	if longURL == "" {
		c.String(http.StatusBadRequest, "enter a url to shorten")
		return
	} else {
		result := dataProccess(longURL)
		if result != "" {
			c.String(http.StatusOK, "Your new link is: http://localhost:8080/"+result)

		} else {
			c.String(http.StatusBadRequest, "invalid URL")
		}
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

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/:id", getReq)
	router.POST("/create", postReq)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run(":8080")

}
