package main

import (
	"github.com/gin-gonic/gin"
)

type ShortenResponse struct {
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
}

// Make this request from the frontend with this: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch#uploading_json_data
// great example: https://www.digitalocean.com/community/tutorials/how-to-use-the-javascript-fetch-api-to-get-data
func testJson(c *gin.Context) {
	// Shorten the URL up here..

	longUrl := "http://google.com"
	shortenedUrl := "http://localhost:8080/ajkhaskjh"
	response := ShortenResponse{
		LongURL:  longUrl,
		ShortURL: shortenedUrl,
	}

	c.JSON(200, response)
}

func main() {

	router := gin.Default()

	router.GET("/test", testJson)

	// router.GET("/:id", getReq)

	// router.GET("/home", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"hello": "world"})
	// })

	router.Run(":8090")
}
