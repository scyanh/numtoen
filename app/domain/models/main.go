package model

import "github.com/gin-gonic/gin"

const (
	StatusSuccess = "ok"
)

type ResponseContext struct {
	Ctx *gin.Context
}

type Response struct {
	Status       string `json:"status"`
	NumInEnglish string `json:"num_in_english"`
}

func (self *ResponseContext) ResponseData(status string, message string) Response {
	response := Response{
		Status:  status,
		NumInEnglish: message,
	}

	return response
}
