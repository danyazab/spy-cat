package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"spy-cat/internal/models"
	"strconv"
)

func (h *MissionHandler) AddTarget(c echo.Context) error {
	var target models.Target
	if err := c.Bind(&target); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.Service.AddTarget(&target); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, target)
}

func (h *MissionHandler) UpdateTarget(c echo.Context) error {
	missionID, _ := strconv.Atoi(c.Param("missionId"))
	targetID, _ := strconv.Atoi(c.Param("targetId"))

	var target models.Target
	if err := c.Bind(&target); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	target.ID = targetID
	target.MissionID = missionID

	if err := h.Service.UpdateTarget(&target); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *MissionHandler) DeleteTarget(c echo.Context) error {
	missionID, _ := strconv.Atoi(c.Param("missionId"))
	targetID, _ := strconv.Atoi(c.Param("targetId"))

	if err := h.Service.DeleteTarget(missionID, targetID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
