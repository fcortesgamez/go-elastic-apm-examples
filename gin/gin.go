package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
	"net/http"
)

func main() {
	engine := gin.New()
	engine.Use(apmgin.Middleware(engine))

	engine.GET("/books/:title/page/:page", serveBookAndPage)

	_ = engine.Run(":8080")
}

func serveBookAndPage(c *gin.Context) {
	title := c.Param("title")
	page := c.Param("page")
	message := fmt.Sprintf("You've requested the book: %s on page %s\n", title, page)
	c.String(http.StatusOK, message)
}
