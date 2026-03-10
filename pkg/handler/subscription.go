package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/effective"
	"github.com/gin-gonic/gin"
)

// @Summary      Создать подписку
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        input  body      effective.Sub  true  "Данные подписки"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  errorResponse
// @Failure      500    {object}  errorResponse
// @Router       /subscriptions/ [post]
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

// @Summary      Получить подписки пользователя
// @Tags         subscriptions
// @Produce      json
// @Param        id    path      string  true   "ID пользователя (UUID)"
// @Param        name  query     string  false  "Фильтр по названию сервиса"
// @Success      200   {array}   effective.Sub
// @Failure      500   {object}  errorResponse
// @Router       /subscriptions/{id} [get]
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

// @Summary      Обновить подписку
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id     path      string                    true  "ID подписки (UUID)"
// @Param        input  body      effective.UpdateSubInput  true  "Поля для обновления (минимум одно)"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  errorResponse
// @Failure      404    {object}  errorResponse
// @Failure      500    {object}  errorResponse
// @Router       /subscriptions/{id} [patch]
func (h *Handler) updateSub(c *gin.Context) {
	subID := c.Param("id")

	var input effective.UpdateSubInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Subscription.UpdateSub(subID, input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newErrorResponse(c, http.StatusNotFound, "subscription not found")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// @Summary      Удалить подписку
// @Tags         subscriptions
// @Produce      json
// @Param        id    path      string  true  "ID пользователя (UUID)"
// @Param        name  query     string  true  "Название сервиса"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  errorResponse
// @Failure      404   {object}  errorResponse
// @Failure      500   {object}  errorResponse
// @Router       /subscriptions/{id} [delete]
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

// @Summary      Суммарная стоимость подписок за период
// @Tags         subscriptions
// @Produce      json
// @Param        user_id       query     string  true   "ID пользователя (UUID)"
// @Param        start_date    query     string  true   "Начало периода (MM-YYYY)"
// @Param        end_date      query     string  true   "Конец периода (MM-YYYY)"
// @Param        service_name  query     string  false  "Фильтр по названию сервиса"
// @Success      200           {object}  map[string]int
// @Failure      400           {object}  errorResponse
// @Failure      500           {object}  errorResponse
// @Router       /subscriptions/total-cost [get]
func (h *Handler) getCost(c *gin.Context) {
	filter := effective.CostFilter{
		UserID:    c.Query("user_id"),
		StartDate: c.Query("start_date"),
		EndDate:   c.Query("end_date"),
	}

	if filter.UserID == "" || filter.StartDate == "" || filter.EndDate == "" {
		newErrorResponse(c, http.StatusBadRequest, "user_id, start_date and end_date are required")
		return
	}

	if name := c.Query("service_name"); name != "" {
		filter.ServiceName = &name
	}

	total, err := h.services.Subscription.GetTotalCost(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_cost": total})
}
