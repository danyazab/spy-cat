package repository

import (
	"errors"
	"spy-cat/internal/models"
)

func (r *missionRepository) AddTarget(target *models.Target) error {
	var complete bool
	query := `SELECT complete FROM missions WHERE id=$1`
	err := r.db.QueryRow(query, target.MissionID).Scan(&complete)
	if err != nil {
		return err
	}
	if complete {
		return errors.New("cannot add target to a completed mission")
	}

	query = `INSERT INTO targets (mission_id, name, country, notes, complete) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.db.Exec(query, target.MissionID, target.Name, target.Country, target.Notes, target.Complete)
	return err
}

func (r *missionRepository) UpdateTarget(target *models.Target) error {
	var complete bool
	query := `SELECT complete FROM targets WHERE id=$1`
	err := r.db.QueryRow(query, target.ID).Scan(&complete)
	if err != nil {
		return err
	}
	if complete {
		return errors.New("cannot update a completed target")
	}

	query = `UPDATE targets SET name=$1, country=$2, notes=$3, complete=$4 WHERE id=$5`
	_, err = r.db.Exec(query, target.Name, target.Country, target.Notes, target.Complete, target.ID)
	return err
}

func (r *missionRepository) DeleteTarget(missionID, targetID int) error {
	var complete bool
	query := `SELECT complete FROM targets WHERE id=$1`
	err := r.db.QueryRow(query, targetID).Scan(&complete)
	if err != nil {
		return err
	}
	if complete {
		return errors.New("cannot delete a completed target")
	}

	query = `DELETE FROM targets WHERE id=$1 AND mission_id=$2`
	_, err = r.db.Exec(query, targetID, missionID)
	return err
}

func (r *missionRepository) GetTargetByID(targetID int) (*models.Target, error) {
	var target models.Target
	err := r.db.QueryRow("SELECT id, mission_id, name, country, notes, complete FROM targets WHERE id = $1", targetID).Scan(
		&target.ID, &target.MissionID, &target.Name, &target.Country, &target.Notes, &target.Complete)
	if err != nil {
		return nil, err
	}
	return &target, nil
}

func (r *missionRepository) UpdateNotes(targetID int, notes string) error {
	var complete bool
	query := `SELECT complete FROM targets WHERE id=$1`
	err := r.db.QueryRow(query, targetID).Scan(&complete)
	if err != nil {
		return err
	}
	if complete {
		return errors.New("cannot update notes for a completed target")
	}

	query = `UPDATE targets SET notes=$1 WHERE id=$2`
	_, err = r.db.Exec(query, notes, targetID)
	return err
}
