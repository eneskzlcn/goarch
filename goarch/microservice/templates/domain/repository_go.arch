package domain

import "database/sql"

type DB interface {
	QueryRow(query string, args ...any) *sql.Row
}
type Repository struct {
	logger Logger
	db     DB
}

func NewRepository(db DB, logger Logger) *Repository {
	if db == nil || logger == nil {
		return nil
	}
	return &Repository{logger: logger, db: db}
}

func (r *Repository) Example() (string, error) {
	query := `SELECT example from examples`
	row := r.db.QueryRow(query)
	var example string
	if err := row.Scan(&example); err != nil {
		return "", err
	}
	return example, nil
}
