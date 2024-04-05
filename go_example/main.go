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
            "date":   time.Now().Format("2024-04-20"),
        })
    })

    router.Run(":8040")
}
