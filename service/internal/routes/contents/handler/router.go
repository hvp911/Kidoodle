package handler

import (
	"context"

	"github.com/gin-gonic/gin"

	"content/internal/routes/contents"
)

// Reader define reader service interface
type Reader interface {
	GetContent(ctx context.Context) ([]contents.Content, error)
}

// Handler endpoint collection
type Handler struct {
	reader Reader
}

// New creates the users transport layer from a service.
func New(readerSvc Reader) Handler {
	return Handler{reader: readerSvc}
}

// SetupRoutes definition of exposed route
func (h Handler) SetupRoutes(users gin.IRouter) {
	v1 := users.Group("/content")
	v1.GET("", h.GetContent)
}
