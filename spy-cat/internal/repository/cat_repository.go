package repository

import (
	"spy-cat/internal/database"
	"spy-cat/internal/models"
)

type CatRepository interface {
	Create(cat *models.Cat) error
	GetByID(id int) (*models.Cat, error)
	UpdateSalary(id int, salary int) error
	Delete(id int) error
	List() ([]models.Cat, error)
}

type catRepository struct {
	db database.Database
}

func NewCatRepository(db database.Database) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Create(cat *models.Cat) error {
	query := `INSERT INTO cats (name, years_of_experience, breed, salary) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, cat.Name, cat.YearsOfExperience, cat.Breed, cat.Salary)
	return err
}

func (r *catRepository) GetByID(id int) (*models.Cat, error) {
	var cat models.Cat
	query := `SELECT id, name, years_of_experience, breed, salary FROM cats WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&cat.ID, &cat.Name, &cat.YearsOfExperience, &cat.Breed, &cat.Salary)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *catRepository) UpdateSalary(id int, salary int) error {
	query := `UPDATE cats SET salary=$1 WHERE id=$2`
	_, err := r.db.Exec(query, salary, id)
	return err
}

func (r *catRepository) Delete(id int) error {
	query := `DELETE FROM cats WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *catRepository) List() ([]models.Cat, error) {
	var cats []models.Cat
	query := `SELECT id, name, years_of_experience, breed, salary FROM cats`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Cat
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.YearsOfExperience, &cat.Breed, &cat.Salary); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}

	return cats, nil
}
