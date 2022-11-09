package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	meme "github.com/mihailco/memessenger"
)

func (h *Handler) createConversation(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input meme.ConversationStruct
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.CreatorId = userId
	id, err := h.services.Conversation.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []meme.ConversationStruct `json:"data"`
}
type getAllUsersResponse struct {
	Data []meme.Userlist `json:"data"`
}

func (h *Handler) getAllConversations(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.Conversation.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) UpdateById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input meme.ConversationStruct
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Id = itemId
	err = h.services.Conversation.UpdateById(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) DeleteById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteById(userId, itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) GetAllUsersAtConv(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if bl := h.services.IsAMember(userId, itemId); bl {
		NewErrorResponse(c, http.StatusTeapot, err.Error())
		return
	}

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	users, err := h.services.GetAllUsersAtConv(itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}

func (h *Handler) AddUserAtConv(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	useridToAdd, err := strconv.Atoi(c.Query("userid"))
	convId, err := strconv.Atoi(c.Query("convid"))

	if bl := h.services.IsAMember(userId, convId); !bl {
		NewErrorResponse(c, http.StatusTeapot, err.Error())
		return
	}

	if err = h.services.AddUser(useridToAdd, convId); err != nil {

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
