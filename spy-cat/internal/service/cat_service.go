package service

import (
	"spy-cat/internal/models"
	"spy-cat/internal/repository"
)

type CatService interface {
	CreateCat(cat *models.Cat) error
	GetCatByID(id int) (*models.Cat, error)
	UpdateCatSalary(id int, salary int) error
	DeleteCat(id int) error
	ListCats() ([]models.Cat, error)
}

type catService struct {
	repo repository.CatRepository
}

func NewCatService(repo repository.CatRepository) CatService {
	return &catService{repo: repo}
}

func (s *catService) CreateCat(cat *models.Cat) error {
	return s.repo.Create(cat)
}

func (s *catService) GetCatByID(id int) (*models.Cat, error) {
	return s.repo.GetByID(id)
}

func (s *catService) UpdateCatSalary(id int, salary int) error {
	return s.repo.UpdateSalary(id, salary)
}

func (s *catService) DeleteCat(id int) error {
	return s.repo.Delete(id)
}

func (s *catService) ListCats() ([]models.Cat, error) {
	return s.repo.List()
}
