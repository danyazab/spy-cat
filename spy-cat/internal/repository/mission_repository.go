package repository

import (
	"errors"
	"spy-cat/internal/database"
	"spy-cat/internal/models"
)

type MissionRepository interface {
	Create(mission *models.Mission) error
	GetByID(id int) (*models.Mission, error)
	UpdateComplete(id int, complete bool) error
	Delete(id int) error
	List() ([]models.Mission, error)

	AddTarget(target *models.Target) error
	UpdateTarget(target *models.Target) error
	DeleteTarget(missionID, targetID int) error
	GetTargetByID(targetID int) (*models.Target, error)

	UpdateNotes(targetID int, notes string) error
	AssignCatToMission(missionID, catID int) error
}

type missionRepository struct {
	db database.Database
}

func NewMissionRepository(db database.Database) MissionRepository {
	return &missionRepository{db: db}
}

func (mq *missionRepository) Create(mission *models.Mission) error {
	if len(mission.Targets) < 1 || len(mission.Targets) > 3 {
		return errors.New("a mission must have between 1 and 3 targets")
	}

	query := `INSERT INTO missions (cat_id, complete) VALUES ($1, $2) RETURNING id`
	err := mq.db.QueryRow(query, mission.CatID, mission.Complete).Scan(&mission.ID)
	if err != nil {
		return err
	}

	for _, target := range mission.Targets {
		query := `INSERT INTO targets (mission_id, name, country, notes, complete) VALUES ($1, $2, $3, $4, $5)`
		_, err = mq.db.Exec(query, mission.ID, target.Name, target.Country, target.Notes, target.Complete)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mq *missionRepository) GetByID(id int) (*models.Mission, error) {
	var mission models.Mission
	query := `SELECT id, cat_id, complete FROM missions WHERE id=$1`
	err := mq.db.QueryRow(query, id).Scan(&mission.ID, &mission.CatID, &mission.Complete)
	if err != nil {
		return nil, err
	}

	query = `SELECT id, mission_id, name, country, notes, complete FROM targets WHERE mission_id=$1`
	rows, err := mq.db.Query(query, mission.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var target models.Target
		if err := rows.Scan(&target.ID, &target.MissionID, &target.Name, &target.Country, &target.Notes, &target.Complete); err != nil {
			return nil, err
		}
		mission.Targets = append(mission.Targets, target)
	}

	return &mission, nil
}

func (mq *missionRepository) UpdateComplete(id int, complete bool) error {
	query := `UPDATE missions SET complete=$1 WHERE id=$2`
	_, err := mq.db.Exec(query, complete, id)
	return err
}

func (mq *missionRepository) Delete(id int) error {
	query := `DELETE FROM missions WHERE id=$1 AND cat_id IS NULL`
	_, err := mq.db.Exec(query, id)
	return err
}

func (mq *missionRepository) List() ([]models.Mission, error) {
	var missions []models.Mission
	query := `SELECT id, cat_id, complete FROM missions`
	rows, err := mq.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mission models.Mission
		if err := rows.Scan(&mission.ID, &mission.CatID, &mission.Complete); err != nil {
			return nil, err
		}

		query = `SELECT id, mission_id, name, country, notes, complete FROM targets WHERE mission_id=$1`
		targetRows, err := mq.db.Query(query, mission.ID)
		if err != nil {
			return nil, err
		}
		defer targetRows.Close()

		for targetRows.Next() {
			var target models.Target
			if err := targetRows.Scan(&target.ID, &target.MissionID, &target.Name, &target.Country, &target.Notes, &target.Complete); err != nil {
				return nil, err
			}
			mission.Targets = append(mission.Targets, target)
		}

		missions = append(missions, mission)
	}

	return missions, nil
}

func (mq *missionRepository) AssignCatToMission(missionID, catID int) error {
	_, err := mq.db.Exec("UPDATE missions SET cat_id = $1 WHERE id = $2", catID, missionID)
	return err
}
