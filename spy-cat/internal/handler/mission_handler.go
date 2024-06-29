package handlers

import (
	"net/http"
	"spy-cat/internal/models"
	"spy-cat/internal/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MissionHandler struct {
	Service service.MissionService
}

func NewMissionHandler(s service.MissionService) *MissionHandler {
	return &MissionHandler{Service: s}
}

func (h *MissionHandler) CreateMission(c echo.Context) error {
	var mission models.Mission
	if err := c.Bind(&mission); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.Service.CreateMission(&mission); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, mission)
}

func (h *MissionHandler) GetMission(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	mission, err := h.Service.GetMissionByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, mission)
}

func (h *MissionHandler) UpdateMission(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var mission models.Mission
	if err := c.Bind(&mission); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.Service.UpdateMissionComplete(id, mission.Complete); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *MissionHandler) DeleteMission(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteMission(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *MissionHandler) ListMissions(c echo.Context) error {
	missions, err := h.Service.ListMissions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, missions)
}

func (h *MissionHandler) AssignCatToMission(c echo.Context) error {
	missionID, _ := strconv.Atoi(c.Param("id"))
	var request struct {
		CatID int `json:"cat_id"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	mission, err := h.Service.GetMissionByID(missionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if mission.Complete {
		return c.JSON(http.StatusForbidden, "Cannot assign a cat to a completed mission")
	}

	if err := h.Service.AssignCatToMission(missionID, request.CatID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
