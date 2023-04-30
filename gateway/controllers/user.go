package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hidekimva/golang/databases/models"
	"github.com/hidekimva/golang/gateway/services"
	"github.com/hidekimva/golang/rabbitmq"
)

type CreateSucessful struct {
	Message  string
	ID       string
	Name     string
	Username string
	Email    string
}

func CreateUser(c *gin.Context) {
	services.LoadEnv()

	amqpUrl := os.Getenv("AMQPURL")
	amqpUser := os.Getenv("AMQPUSER")
	amqpPass := os.Getenv("AMQPPASSWORD")

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rabbitmq.SendMessage(amqpUser, amqpPass, amqpUrl, true, "create-user", "create-user-r", user)

	msg := rabbitmq.ConsumeMessageQueue(amqpUser, amqpPass, amqpUrl, "create-user-r")

	err := json.Unmarshal(msg.Body, &user)

	if err != nil {
		log.Fatalf("Falha ao decodificar objeto JSON: %v", err)
	}

	data := CreateSucessful{
		Message:  "User created successfully",
		ID:       user.ID,
		Name:     user.Name,
		Username: user.UserName,
		Email:    user.Email,
	}

	c.JSON(200, data)

}
