package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
	"github.com/naumovrus/backend-trainee-asignment/interntal/repository"
)

func (h *Handler) createSegment(c *gin.Context) {
	var input ent.Segment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Segment.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type getAllIdsResponse struct {
	Ids []int `json:"ids"`
}

func (h *Handler) addUserSegment(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid userId param")
		return
	}
	var segments repository.SegmentRequest // fix and add segmentReq in pkg or smth

	if err := c.BindJSON(&segments); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ids, err := h.services.Segment.AddUserSegment(userId, segments)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}
	c.JSON(http.StatusCreated, getAllIdsResponse{
		Ids: ids,
	})
}

func (h *Handler) deleteUserSegment(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid userId param")
		return
	}

	var segments repository.SegmentRequest // fix and add segmentReq in pkg or smth

	if err := c.BindJSON(&segments); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.DeleteUserSegment(userId, segments)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteSegment(c *gin.Context) {
	var segment ent.Segment
	if err := c.BindJSON(&segment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.DeleteSegment(segment)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
