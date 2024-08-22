package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorBase error message struct returned back as a response upon request failures
type ErrorBase struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// HTTPError Error struct to be used by handler. This struct takes care of setting error codes and messages correctly before hand.
type HTTPError struct {
	HTTPCode  int
	ErrorBase ErrorBase
}

// Error returns error message as a string
func (h *HTTPError) Error() string {
	return h.ErrorBase.Message
}

// GetError re-builds error message if error message needs to be re built using args
func (h *HTTPError) GetError(args ...interface{}) *HTTPError {
	return &HTTPError{
		HTTPCode: h.HTTPCode,
		ErrorBase: ErrorBase{
			Message: fmt.Sprintf(h.ErrorBase.Message, args...),
			Code:    h.ErrorBase.Code,
		},
	}
}

// newHTTPError helper function to prepare error objects to be used by handlers
func newHTTPError(httpCode int, message string, errorCode int) *HTTPError {
	return &HTTPError{
		HTTPCode: httpCode,
		ErrorBase: ErrorBase{
			Message: message,
			Code:    errorCode,
		},
	}
}

// Define error responses as a variable
var (
	InputError    = newHTTPError(http.StatusBadRequest, "invalid query parameter: %s", 4000)
	InputIDError  = newHTTPError(http.StatusBadRequest, "invalid id parameter: %s", 4000)
	ServerError   = newHTTPError(http.StatusInternalServerError, "internal server error", 5000)
	DatabaseError = newHTTPError(http.StatusInternalServerError, "database error: %s", 5002)
)

// WriteSuccessResponse writes success response to the response in JSON format
func WriteSuccessResponse(ctx *gin.Context, httpStatusCode int, response interface{}) {
	ctx.JSON(httpStatusCode, response)
}

// WriteErrorResponse writes error response to the response in JSON format. This function takes care of converting unknown errors to generic error format.
func WriteErrorResponse(ctx *gin.Context, err error) {
	writer := ctx.Writer
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var httpError *HTTPError
	ok := errors.As(err, &httpError)
	if !ok {
		writer.WriteHeader(http.StatusInternalServerError)
		encodeMessage(writer, ServerError.ErrorBase)
		return
	}
	writer.WriteHeader(httpError.HTTPCode)
	encodeMessage(writer, httpError.ErrorBase)
}

func encodeMessage(writer gin.ResponseWriter, message interface{}) {
	ee := json.NewEncoder(writer).Encode(message)
	if ee != nil {
		log.Fatalf("encode message error: %v", ee)
	}
}
