package handler

import (
	"net/http"

	"github.com/effective"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createSub(c *gin.Context) {
	var input effective.Sub

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Subscription.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getSub(c *gin.Context) {

}

func (h *Handler) getAllSubs(c *gin.Context) {

}

func (h *Handler) updateSub(c *gin.Context) {

}

func (h *Handler) deleteSub(c *gin.Context) {

}

func (h *Handler) deleteAllSubs(c *gin.Context) {

}

func (h *Handler) getCost(c *gin.Context) {

}

func (h *Handler) getUserInfo(c *gin.Context) {

}
