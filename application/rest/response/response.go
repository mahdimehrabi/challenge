package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BaseResponse is a custom response structure.
type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Respond is a general function to generate responses with a status code, message, and data.
func Respond(c *gin.Context, status int, message string, data interface{}) {
	response := BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.JSON(status, response)
}

// RespondOK generates a response with a 200 OK status.
func RespondOK(c *gin.Context, message string, data interface{}) {
	Respond(c, http.StatusOK, message, data)
}

// RespondBadRequest generates a response with a 400 Bad Request status.
func RespondBadRequest(c *gin.Context, data interface{}) {
	Respond(c, http.StatusBadRequest, "bad request", data)
}

// RespondInternalServerError generates a response with a 500 Internal Server Error status.
func RespondInternalServerError(c *gin.Context, data interface{}) {
	Respond(c, http.StatusInternalServerError, "internal server error", data)
}
