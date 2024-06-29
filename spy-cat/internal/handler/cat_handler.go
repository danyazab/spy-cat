package handlers

import (
	"net/http"
	"spy-cat/internal/models"
	"spy-cat/internal/service"
	validator "spy-cat/pkg/validator_breed"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CatHandler struct {
	Service service.CatService
}

func NewCatHandler(s service.CatService) *CatHandler {
	return &CatHandler{Service: s}
}

func (h *CatHandler) CreateCat(c echo.Context) error {
	cat := new(models.Cat)
	if err := c.Bind(cat); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate breed using TheCatAPI
	if !validator.ValidateBreed(cat.Breed) {
		return c.JSON(http.StatusBadRequest, "Invalid breed")
	}

	err := h.Service.CreateCat(cat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, cat)
}

func (h *CatHandler) GetCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, err := h.Service.GetCatByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) UpdateCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cat models.Cat
	if err := c.Bind(&cat); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.Service.UpdateCatSalary(id, cat.Salary); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CatHandler) DeleteCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteCat(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *CatHandler) ListCats(c echo.Context) error {
	cats, err := h.Service.ListCats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cats)
}
