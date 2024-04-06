package main

import (
    "github.com/gin-gonic/gin"
    "time"
)

const (
	stdZeroMonth = "01"
	stdZeroDay   = "02"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
	router.Static("/css", "./css")
    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "haveto": []string{"Michael", "Dwight", "Jim", "Pam", "Angela", "Oscar", "Kevin"},
            "did":    []string{"Stanley", "Phyllis", "Andy", "Robert", "Ryan", "Kelly", "Toby"},
            "date":   time.Now().Format("2006-01-02"),
            "month":   time.Now().Format(stdZeroMonth),
            "day":   time.Now().Format(stdZeroDay),
        })
    })
    router.GET("/list_did.html", func(c *gin.Context) {
        c.HTML(200, "list_did.html", gin.H{
            "did":    []string{"Stanley", "Phyllis", "Andy", "Robert", "Ryan", "Kelly", "Toby"},
        })
    })
    router.GET("/list_have_to.html", func(c *gin.Context) {
        c.HTML(200, "list_did.html", gin.H{
            "haveto":    []string{"Stanley", "Phyllis", "Andy", "Robert", "Ryan", "Kelly", "Toby"},
        })
    })
    router.Run(":8040")
}
