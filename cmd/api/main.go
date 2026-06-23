package main

import (
	"github.com/gin-gonic/gin"
	"GoAssistAutoBackend/internal/handlers"
	//"GoAssistAutoBackend/internal/config"
	// "GoAssistAutoBackend/internal/models" importacion de los modelos
)


func main(){
	//Se incia la base de datos
	//db := config.InitDB()
	// aqui se migran los modelos de Internal models
	// db.AutoMigrate(&models.User) user es un ejemplo
	//Aqui se generan el enrutamiento
	router := gin.Default()
	router.GET("/hola/:name",handlers.Hello)
	router.POST("/alumnos", handlers.CrearAlumno)
	router.Run()
}
