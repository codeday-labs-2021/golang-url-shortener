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

var dataArray map[string]string

func genID(length int) string {

	id, err := gonanoid.New(length)
	if err != nil {
		panic(err)
	}
	return id
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
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
					// if true {
					dataArray[currID] = inputURL
					result, err := json.Marshal(dataArray)
					if err != nil {
						log.Println(err)
					}
					ioutil.WriteFile("urlmap.json", result, 0644)

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

	fmt.Println(id)
	_, exists := dataArray[id]
	if exists {
		c.Redirect(http.StatusMovedPermanently, dataArray[id])
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

	jsonFile, err := os.Open("urlmap.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened url.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &dataArray); err != nil {
		log.Println(err)
	}

	dataProccess()

	router := routerSetup()

	router.Run(":8090")

}
// example.com/ISq
// example.com/