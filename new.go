package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// generate ID

func genID(length int) string {

	id, err := gonanoid.New(length)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated id: %s\n", id)
	return id
}

// check if url
func isUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// make database
type urldatabase struct {
	UrlID   string `json:"urlID"`
	LongURL string `json:"longURL"`
}

// validate then store
func valjsonstore(URLinput string) {
	M := make(map[string]string)
	//sc := bufio.NewScanner(os.Stdin)

	//generage new ID and store
	idLength := 5
	//store URL along with currID
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

	if isUrl(URLinput) {
		for {
			currID := genID(idLength)
			if _, ok := M[currID]; !ok {
				M[currID] = URLinput
				dataArray = append(dataArray, urldatabase{UrlID: currID, LongURL: URLinput})
				for _, v := range dataArray {
					fmt.Println(v)
				}
				break
			} else {
				idLength += 1
			}
		}
	} else {
		fmt.Print("Url not valid")
	}

}

//yoink data from json

//redirects

func setupRouter() *gin.Engine {
	router := gin.Default()
	// put get and post requests in here
	return router
}

func main() {
	var URLinput string
	fmt.Print("Plz paste url")
	fmt.Scan(&URLinput)
	//valjsonstore(URLinput)
	valjsonstore(URLinput)
}
