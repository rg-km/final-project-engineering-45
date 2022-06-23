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
	prodi.created_at,
	prodi.deskripsi,
	prodi.karakter,
	prodi.mata_kuliah,
	prodi.prospek
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
		&prodi.Deskripsi,
		&prodi.Karakter,
		&prodi.MataKuliah,
		&prodi.Prospek,
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
		prodi.created_at,
		prodi.deskripsi,
		prodi.karakter,
		prodi.mata_kuliah,
		prodi.prospek
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
			&prodiRow.Deskripsi,
			&prodiRow.Karakter,
			&prodiRow.MataKuliah,
			&prodiRow.Prospek,
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
	prodi.created_at,
	prodi.deskripsi,
	prodi.karakter,
	prodi.mata_kuliah,
	prodi.prospek
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
		&prodi.Deskripsi,
		&prodi.Karakter,
		&prodi.MataKuliah,
		&prodi.Prospek,
	)
	if err != nil {
		return prodi, err
	}

	return prodi, nil
}

func (p *ProdiRepository) FetchProdiByFakultasName(fakultasName string) ([]ProgramStudi, error) {
	var prodi []ProgramStudi

	sqlStatement := `SELECT 
	prodi.id,
	prodi.prodi_name,
	prodi.fakultas_id,
	fakultas.fakultas_name,
	prodi.created_at,
	prodi.deskripsi,
	prodi.karakter,
	prodi.mata_kuliah,
	prodi.prospek
	FROM program_studi prodi
	INNER JOIN fakultas ON prodi.fakultas_id = fakultas.id
	WHERE fakultas.fakultas_name = ?`

	rows, err := p.db.Query(sqlStatement, fakultasName)
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
			&prodiRow.Deskripsi,
			&prodiRow.Karakter,
			&prodiRow.MataKuliah,
			&prodiRow.Prospek,
		)

		if err != nil {
			return prodi, err
		}

		prodi = append(prodi, prodiRow)
	}

	return prodi, nil
}

func (p *ProdiRepository) FetchProdiAndFakultasName(fakultasName string, prodiName string) (ProgramStudi, error) {
	var prodi ProgramStudi

	sqlStatement := `SELECT
	prodi.id,
	prodi.prodi_name,
	prodi.fakultas_id,
	fakultas.fakultas_name,
	prodi.created_at,
	prodi.deskripsi,
	prodi.karakter,
	prodi.mata_kuliah,
	prodi.prospek
	FROM program_studi prodi
	INNER JOIN fakultas ON prodi.fakultas_id = fakultas.id
	WHERE fakultas.fakultas_name = ? AND prodi.prodi_name = ?`

	rows := p.db.QueryRow(sqlStatement, fakultasName, prodiName)
	err := rows.Scan(
		&prodi.ID,
		&prodi.ProdiName,
		&prodi.FakultasID,
		&prodi.FakultasName,
		&prodi.CreatedAt,
		&prodi.Deskripsi,
		&prodi.Karakter,
		&prodi.MataKuliah,
		&prodi.Prospek,
	)

	if err != nil {
		return prodi, err
	}

	return prodi, nil
}

func (p *ProdiRepository) InsertProdi(prodiName string, fakultasName string, deskripsi string, karakter string, mataKuliah string, prospek string) error {
	sqlStatement := `INSERT INTO program_studi (id, prodi_name, fakultas_id, created_at, deskripsi, karakter, mata_kuliah, prospek) 
	VALUES (
		(SELECT id FROM program_studi WHERE fakultas_id = (SELECT id FROM fakultas WHERE fakultas_name = ?) ORDER BY id DESC LIMIT 1) + 1, 
		?, 
		(SELECT id FROM fakultas WHERE fakultas_name = ?), 
		CURRENT_TIMESTAMP, ?, ?, ?, ?
	)`

	_, err := p.db.Exec(sqlStatement, fakultasName, prodiName, fakultasName, deskripsi, karakter, mataKuliah, prospek)
	if err != nil {
		return err
	}

	return nil
}
