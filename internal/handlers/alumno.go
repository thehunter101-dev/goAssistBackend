package handlers

import (
	"net/http"

	"GoAssistAutoBackend/internal/config"
	"GoAssistAutoBackend/internal/models"
	"github.com/gin-gonic/gin"
)

func CrearAlumno(c *gin.Context) {

	var alumno models.Alumno

	if err := c.ShouldBindJSON(&alumno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	alumno.Estado = "Ausente"

	if err := config.DB.Create(&alumno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, alumno)
}