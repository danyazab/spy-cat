package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"spy-cat/internal/models"
	"strconv"
)

func (h *MissionHandler) UpdateNotes(c echo.Context) error {
	targetID, _ := strconv.Atoi(c.Param("targetId"))
	var target models.Target
	if err := c.Bind(&target); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	existingTarget, err := h.Service.GetTargetByID(target.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	mission, err := h.Service.GetMissionByID(existingTarget.MissionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if existingTarget.Complete || mission.Complete {
		return c.JSON(http.StatusForbidden, "Notes cannot be updated if the target or mission is completed")
	}

	if err := h.Service.UpdateNotes(targetID, target.Notes); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, target)
}
