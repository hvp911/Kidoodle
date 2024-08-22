package handler

import (
	"content/internal"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"

	_ "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// GetContent contents handler.
func (h Handler) GetContent(ctx *gin.Context) {
	requestContext := ctx.Request.Context()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Printf("error getting logger: %v", err)
		internal.WriteErrorResponse(ctx, internal.ServerError.GetError(err.Error()))
		return
	}

	// retrieve contents from reader service using request context
	contentList, err := h.reader.GetContent(requestContext)
	if err != nil {
		logger.Error(fmt.Sprintf("error fetching contents: %v", err))
		internal.WriteErrorResponse(ctx, internal.DatabaseError.GetError(err.Error()))
		return
	}

	logger.Info("successfully fetched content list")
	internal.WriteSuccessResponse(ctx, http.StatusOK, contentList)
}
