package repository

import (
	"database/sql"
)

type FakultasRepository struct {
	db *sql.DB
}

func NewFakultasRepository(db *sql.DB) *FakultasRepository {
	return &FakultasRepository{db: db}
}

func (f *FakultasRepository) FetchFakultas() ([]Fakultas, error) {
	sqlStatement := `SELECT * from fakultas`
	var fakultas []Fakultas
	rows, err := f.db.Query(sqlStatement)
	if err != nil {
		return fakultas, err
	}

	for rows.Next() {
		var fakultasRow Fakultas
		err = rows.Scan(&fakultasRow.ID, &fakultasRow.FakultasName, &fakultasRow.CreatedAt)
		if err != nil {
			return nil, err
		}

		fakultas = append(fakultas, fakultasRow)
	}

	return fakultas, nil
}

func (f *FakultasRepository) FetchFakultasByName(fakultasName string) (Fakultas, error) {
	sqlStatement := `SELECT * from fakultas WHERE fakultas_name = ?`
	var fakultas Fakultas
	rows, err := f.db.Query(sqlStatement, fakultasName)
	if err != nil {
		return fakultas, err
	}

	for rows.Next() {
		err = rows.Scan(&fakultas.ID, &fakultas.FakultasName, &fakultas.CreatedAt)
		if err != nil {
			return fakultas, err
		}
	}

	return fakultas, nil
}
