// Unused code that might be useless for later

/*
r.GET("/test", func(c *gin.Context) {
	c.Request.URL.Path = "/test2"
	r.HandleContext(c)
})

r.GET("/", func(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
})

r.GET(currID, func(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, inputURL)
})

data := urldatabase{
	UrlID:   currID,
	LongURL: inputURL,
}
fmt.Println(data)

file, _ := json.MarshalIndent(data, "", " ")
fmt.Print(string(file))
ioutil.WriteFile("urlmap.json", file, 0644)
*/