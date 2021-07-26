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

var DATA_ARRAY map[string]string
var SERVER_URL string = "http://localhost:8080"

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

func dataProccess(inputURL string) (string, error) {
	idLength := 3
	if isUrl(inputURL) {
		for {
			currID := genID(idLength)
			if _, exists := DATA_ARRAY[currID]; !exists {
				DATA_ARRAY[currID] = inputURL
				result, err := json.Marshal(DATA_ARRAY)
				if err != nil {
					logrus.Errorf("error while marshalling json: %v", err)
					return "", fmt.Errorf("error marshalling json for URL map: %v", err)
				}
				ioutil.WriteFile("urlmap.json", result, 0644)

				return currID, nil
			} else {
				idLength += 1
			}
		}
	}
	return "", fmt.Errorf("input is not valid URL")

}

func getReq(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)
	url, exists := DATA_ARRAY[id]
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
		result, err := dataProccess(longURL)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid URL")
		} else {
			c.String(http.StatusOK, fmt.Sprintf(
				"Your new link is: %s/%s", SERVER_URL, result))
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
	if err := json.Unmarshal(byteValue, &DATA_ARRAY); err != nil {
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
