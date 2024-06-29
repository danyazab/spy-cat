package service

import (
	"spy-cat/internal/models"
	"spy-cat/internal/repository"
)

type MissionService interface {
	CreateMission(cat *models.Mission) error
	GetMissionByID(id int) (*models.Mission, error)
	UpdateMissionComplete(id int, complete bool) error
	DeleteMission(id int) error
	ListMissions() ([]models.Mission, error)

	AddTarget(target *models.Target) error
	UpdateTarget(target *models.Target) error
	DeleteTarget(missionID, targetID int) error
	GetTargetByID(targetID int) (*models.Target, error)

	UpdateNotes(targetID int, notes string) error
	AssignCatToMission(missionID, catID int) error
}

type missionService struct {
	repo repository.MissionRepository
}

func NewMissionService(repo repository.MissionRepository) MissionService {
	return &missionService{repo: repo}
}

func (s *missionService) CreateMission(cat *models.Mission) error {
	return s.repo.Create(cat)
}

func (s *missionService) GetMissionByID(id int) (*models.Mission, error) {
	return s.repo.GetByID(id)
}

func (s *missionService) UpdateMissionComplete(id int, complete bool) error {
	return s.repo.UpdateComplete(id, complete)
}

func (s *missionService) DeleteMission(id int) error {
	return s.repo.Delete(id)
}

func (s *missionService) ListMissions() ([]models.Mission, error) {
	return s.repo.List()
}

func (s *missionService) AddTarget(target *models.Target) error {
	return s.repo.AddTarget(target)
}

func (s *missionService) UpdateTarget(target *models.Target) error {
	return s.repo.UpdateTarget(target)
}

func (s *missionService) DeleteTarget(missionID, targetID int) error {
	return s.repo.DeleteTarget(missionID, targetID)
}

func (s *missionService) UpdateNotes(targetID int, notes string) error {
	return s.repo.UpdateNotes(targetID, notes)
}

func (s *missionService) AssignCatToMission(missionID, catID int) error {
	return s.repo.AssignCatToMission(missionID, catID)
}

func (s *missionService) GetTargetByID(targetID int) (*models.Target, error) {
	return s.repo.GetTargetByID(targetID)
}
