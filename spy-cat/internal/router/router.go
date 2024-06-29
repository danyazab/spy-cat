package router

import (
	"github.com/labstack/echo/v4"
	handlers "spy-cat/internal/handler"
)

func NewRouter(e *echo.Echo, catHandler *handlers.CatHandler, missionHandler *handlers.MissionHandler) {
	// Missions
	e.POST("/missions", missionHandler.CreateMission)
	e.GET("/missions/:id", missionHandler.GetMission)
	e.PUT("/missions/:id", missionHandler.UpdateMission)
	e.DELETE("/missions/:id", missionHandler.DeleteMission)
	e.GET("/missions", missionHandler.ListMissions)
	e.PUT("/missions/:id/assign", missionHandler.AssignCatToMission)

	// Targets
	e.POST("/missions/:id/targets", missionHandler.AddTarget)
	e.PUT("/missions/:missionId/targets/:targetId", missionHandler.UpdateTarget)
	e.DELETE("/missions/:missionId/targets/:targetId", missionHandler.DeleteTarget)
	e.PUT("/targets/:targetId/notes", missionHandler.UpdateNotes)

	// Cats
	e.POST("/cats", catHandler.CreateCat)
	e.GET("/cats/:id", catHandler.GetCat)
	e.PUT("/cats/:id", catHandler.UpdateCat)
	e.DELETE("/cats/:id", catHandler.DeleteCat)
	e.GET("/cats", catHandler.ListCats)

}
