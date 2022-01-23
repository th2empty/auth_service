package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/th2empty/auth-server/pkg/models"
	"net/http"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var input models.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {

}
