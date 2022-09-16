package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	meme "github.com/mihailco/memessenger"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input meme.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err2 := h.services.Authorization.CreateUser(input)

	if err2 != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput
	logrus.Println(c.Request.Header)

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err2 := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err2 != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
