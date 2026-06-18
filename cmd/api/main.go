package main

import (
	"github.com/gin-gonic/gin"
	"GoAssistAutoBackend/internal/handlers"
)


func main(){
	//Aqui se generan el enrutamiento
	router := gin.Default()
	router.GET("/hola/:name",handlers.Hello)

	router.Run()
}
