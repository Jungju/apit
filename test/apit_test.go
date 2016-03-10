package test

import (
	"testing"
	"github.com/jungju/apit"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"fmt"
)

type requestTestBody struct {
	Name string `json:"name"`
	Message string `json:"message"`
}

type responseTestBody struct {
	ServerMessage string `json:"serverMessage"`
}

func TestApit(t *testing.T){

	r := gin.Default()

	r.GET("/home", routeGetHome)
	r.POST("/home", routePostHome)

	apit.SetGin(r)

	resBody := responseTestBody{}
	apit.PerformRequest(t, "GET", "/home", nil, nil, &resBody, 200)
	assert.Equal(t, "Send from server...", resBody.ServerMessage)

	reqBody := requestTestBody{
		Name:"jungju",
		Message:"Hi~~~",
	}
	apit.PerformRequest(t, "POST", "/home", nil, reqBody, &resBody, 200)
	assert.NotNil(t, resBody)
	assert.Equal(t, fmt.Sprintf("Hi %s!! I love you.", reqBody.Name), resBody.ServerMessage)
}

func routeGetHome(c *gin.Context){
	resBody := responseTestBody{
		ServerMessage:"Send from server...",
	}
	c.JSON(200, resBody)
}

func routePostHome(c *gin.Context){
	reqBody := requestTestBody{}
	c.BindJSON(&reqBody)
	resBody := responseTestBody{
		ServerMessage:fmt.Sprintf("Hi %s!! I love you.", reqBody.Name),
	}
	c.JSON(200, resBody)
}