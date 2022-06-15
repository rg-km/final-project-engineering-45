package repository

import "database/sql"

type ProdiRepository struct {
	db *sql.DB
}

func NewProdiRepository(db *sql.DB) *ProdiRepository {
	return &ProdiRepository{db: db}
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
