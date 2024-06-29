package models

type Mission struct {
	ID       int      `json:"id"`
	CatID    int      `json:"cat_id"`
	Complete bool     `json:"complete"`
	Targets  []Target `json:"targets"`
}
