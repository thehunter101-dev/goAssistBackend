package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context){
	name := c.Param("name")
	saludo := fmt.Sprintf("Hola %s, como estas?", name)
	c.JSON(http.StatusOK, gin.H{
		"Saludo" : saludo,
	})
}
