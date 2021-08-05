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

type Shorten struct {
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
}

type Error struct {
	ErrCode    int    `json:"ErrCode"`
	ErrMessage string `json:"ErrMessage"`
}

var noURL = Error{
	ErrCode:    400,
	ErrMessage: "please enter a url to shorten",
}

var invalidURL = Error{
	ErrCode:    400,
	ErrMessage: "invalid URL",
}

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
	var link Shorten
	c.BindJSON(&link)

	if link.LongURL == "" {
		c.JSON(noURL.ErrCode, noURL)
		return
	}
	result, err := dataProccess(link.LongURL)
	if err != nil {
		c.JSON(invalidURL.ErrCode, invalidURL)
		return
	}

	response := Shorten{
		LongURL:  link.LongURL,
		ShortURL: result,
	}

	c.JSON(200, response)
}

func main() {

	jsonFile, err := os.Open("urlmap.json")
	if os.IsNotExist(err) {
		jsonFile, err = os.Create("urlmap.json")
		if err != nil {
			logrus.Errorf("error while creating urlmap.json: %v", err)
			os.Exit(1)
		}
		ioutil.WriteFile("urlmap.json", []byte("{}"), 0644)
	} else if err != nil {
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

	router.Run()

}
