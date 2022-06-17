package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ProdiRepository struct {
	db *sql.DB
}

func NewProdiRepository(db *sql.DB) *ProdiRepository {
	return &ProdiRepository{db: db}
}

//fetch prodi by name
func (p *ProdiRepository) FetchProdiByName(prodiName string) (ProgramStudi, error) {
	sqlStatement := `SELECT 
	prodi.id,
	prodi.prodi_name,
	prodi.fakultas_id,
	fakultas.fakultas_name,
	prodi.created_at
	FROM program_studi prodi
	INNER JOIN fakultas ON prodi.fakultas_id = fakultas.id
	WHERE prodi.prodi_name = ?`

	var prodi ProgramStudi
	rows := p.db.QueryRow(sqlStatement, prodiName)
	err := rows.Scan(
		&prodi.ID,
		&prodi.ProdiName,
		&prodi.FakultasID,
		&prodi.FakultasName,
		&prodi.CreatedAt,
	)
	if err != nil {
		return prodi, err
	}

	return prodi, nil
}

func (p *ProdiRepository) FetchProdi() ([]ProgramStudi, error) {
	var prodi []ProgramStudi

	//join fakultas and program studi
	sqlStatement := `
	SELECT 
		prodi.id,
		prodi.prodi_name,
		prodi.fakultas_id,
		fakultas.fakultas_name,
		prodi.created_at
	FROM program_studi prodi
	INNER JOIN fakultas ON prodi.fakultas_id = fakultas.id
	`
	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return prodi, err
	}

	for rows.Next() {
		var prodiRow ProgramStudi
		err = rows.Scan(
			&prodiRow.ID,
			&prodiRow.ProdiName,
			&prodiRow.FakultasID,
			&prodiRow.FakultasName,
			&prodiRow.CreatedAt,
		)
		if err != nil {
			return prodi, err
		}

		prodi = append(prodi, prodiRow)
	}

	return prodi, nil
}

//fetch prodi by id
func (p *ProdiRepository) FetchProdiByID(id int64) (ProgramStudi, error) {
	var prodi ProgramStudi

	sqlStatement := `SELECT 
	prodi.id,
	prodi.prodi_name,
	prodi.fakultas_id,
	fakultas.fakultas_name,
	prodi.created_at
	FROM program_studi prodi
	INNER JOIN fakultas ON prodi.fakultas_id = fakultas.id
	WHERE prodi.id = ?`

	rows := p.db.QueryRow(sqlStatement, id)
	err := rows.Scan(
		&prodi.ID,
		&prodi.ProdiName,
		&prodi.FakultasID,
		&prodi.FakultasName,
		&prodi.CreatedAt,
	)
	if err != nil {
		return prodi, err
	}

	return prodi, nil
}

func (p *ProdiRepository) InsertProdi(prodiName string, fakultasName string) error {
	sqlStatement := `INSERT INTO program_studi (prodi_name, fakultas_id, created_at)
	VALUES (?, (SELECT id FROM fakultas WHERE fakultas_name = ?), CURRENT_TIMESTAMP)`

	_, err := p.db.Exec(sqlStatement, prodiName, fakultasName)
	if err != nil {
		return err
	}

	return nil
}
