package main

import (
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "haveto": []string{"Michael", "Dwight", "Jim", "Pam", "Angela", "Oscar", "Kevin"},
            "did":    []string{"Stanley", "Phyllis", "Andy", "Robert", "Ryan", "Kelly", "Toby"},
            "date":   time.Now().Format("2006-01-02"),
        })
    })

    router.Run(":8080")
}
