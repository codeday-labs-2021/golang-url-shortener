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

func dataProccess() {
	sc := bufio.NewScanner(os.Stdin)
	idLength := 3
	for {
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
					// if true {
					dataArray = append(dataArray, urldatabase{UrlID: currID, LongURL: inputURL})

					result, err := json.Marshal(dataArray)
					if err != nil {
						log.Println(err)
					}
					ioutil.WriteFile("urlmap.json", result, 0644)
					jsonFile.Close()

					for _, v := range dataArray {
						fmt.Println(v)
					}

					break
					// } else {
					// 	idLength += 1
					// }
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
	fmt.Println(id)
	for i := range dataArray {
		if dataArray[i].UrlID == id {
			fmt.Println("yes")
			c.Redirect(http.StatusMovedPermanently, dataArray[i].LongURL)
			break
		}
	}

}

func routerSetup() *gin.Engine {
	router := gin.Default()

	router.GET("/:id", getReq)

	router.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	return router
}

func main() {
	dataProccess()

	router := routerSetup()

	router.Run(":8090")
}
