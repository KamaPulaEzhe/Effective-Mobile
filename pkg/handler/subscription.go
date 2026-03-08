package handler

import (
	"database/sql"
	"errors"
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
	id := c.Param("id")
	name := c.Query("name")
	if name != "" {
		sub, err := h.services.GetSub(id, name)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, sub)
		return
	}

	subs, err := h.services.GetAllSubs(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, subs)

}

// func (h *Handler) getAllSubs(c *gin.Context) {
// 	// id := c.Param("id")
// 	// subs, err := h.services.

// 	return
// }

func (h *Handler) updateSub(c *gin.Context) {
}

func (h *Handler) deleteSub(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")

	if name == "" {
		newErrorResponse(c, http.StatusBadRequest, "name query parameter is required")
		return
	}

	err := h.services.DeleteSub(id, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newErrorResponse(c, http.StatusNotFound, "subscription not found")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *Handler) deleteAllSubs(c *gin.Context) {
}

func (h *Handler) getCost(c *gin.Context) {
}

func (h *Handler) getUserInfo(c *gin.Context) {
}
