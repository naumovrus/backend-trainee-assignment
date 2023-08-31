package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naumovrus/backend-trainee-asignment/interntal/entities"
	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
)

func (h *Handler) createUser(c *gin.Context) {
	var input ent.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.User.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type getAllSegmentsResponse struct {
	Data []entities.Segment `json:"segments"`
}

func (h *Handler) getUserSegments(c *gin.Context) {
	var segments []ent.Segment
	userId, err := strconv.Atoi(c.Param("userId"))
	strconv.Atoi(c.Param("userId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid userId param")
		return
	}
	segments, err = h.services.User.GetUserSegments(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllSegmentsResponse{
		Data: segments,
	})
}
