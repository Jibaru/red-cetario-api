package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"redcetarioapi/models"
)

func Register(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "register called", "path", c.Request.URL.Path)

	var input struct {
		Nombre            string `json:"nombre" binding:"required"`
		ApePaterno        string `json:"ape_paterno" binding:"required"`
		ApeMaterno        string `json:"ape_materno" binding:"required"`
		CorreoElectronico string `json:"correo_electronico" binding:"required,email"`
		Contrasenia       string `json:"contrasenia" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "mensaje": err.Error()})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Contrasenia), bcrypt.DefaultCost)
	cliente := models.Cliente{
		Nombre:            input.Nombre,
		ApePaterno:        input.ApePaterno,
		ApeMaterno:        input.ApeMaterno,
		CorreoElectronico: input.CorreoElectronico,
		Contrasenia:       string(hash),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	DB.Create(&cliente)

	logger.InfoContext(c, "client registered", "client_id", cliente.ID)

	c.JSON(http.StatusOK, gin.H{"ok": true, "cliente": cliente})
}

func Login(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "login called", "path", c.Request.URL.Path)

	var input struct {
		CorreoElectronico string `json:"correo_electronico" binding:"required,email"`
		Contrasenia       string `json:"contrasenia" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "mensaje": err.Error()})
		return
	}
	var cliente models.Cliente
	if err := DB.Where("correo_electronico = ?", input.CorreoElectronico).First(&cliente).Error; err != nil {
		logger.ErrorContext(c, "login failed", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "mensaje": "Correo o contraseña inválidas"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(cliente.Contrasenia), []byte(input.Contrasenia)) != nil {
		logger.ErrorContext(c, "login failed", "error", "invalid password")
		c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "mensaje": "Correo o contraseña inválidas"})
		return
	}

	logger.InfoContext(c, "client logged in", "client_id", cliente.ID)

	c.JSON(http.StatusOK, gin.H{"ok": true, "cliente": cliente})
}
