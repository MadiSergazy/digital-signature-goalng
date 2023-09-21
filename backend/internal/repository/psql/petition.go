package psql

import (
	"context"

	"go.uber.org/zap"

	"mado/internal/core/petition"
	"mado/pkg/database/postgres"
)

type PetitionRepository struct {
	db     *postgres.Postgres
	logger *zap.Logger
}

// PetitionRepository creates a new UserRepository.
func NewPetitionRepository(db *postgres.Postgres, logger *zap.Logger) SurveyrRepository {
	return SurveyrRepository{
		db:     db,
		logger: logger,
	}
}

// todo create table for this and think about numering documents IDEA: just before create new file get id for the next row in postgre
func (p PetitionRepository) Save(ctx context.Context, dto *petition.PetitionData) (*petition.PetitionData, error) {
	// Insert the PDF data into the database
	_, err := p.db.Pool.Exec(ctx, `
	  INSERT INTO petition_pdf_files (file_name, file_data, creation_date)
	  VALUES ($1, $2, $3)`,
		dto.FileName, dto.PdfData, dto.CreationDate)
	if err != nil {
		p.logger.Error("Failed to insert PDF data into the database: ", zap.Error(err))
		return nil, err
	}

	return dto, nil
}

// We use the nextval function with the sequence name petitions_id_seq to retrieve the next available ID.
func (p PetitionRepository) GetNextID(ctx context.Context) (*int, error) {
	query := "SELECT nextval('petitions_id_seq')"

	var id int
	if err := p.db.Pool.QueryRow(ctx, query).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

// CREATE TABLE petitions (
//     id SERIAL PRIMARY KEY,
//     file_name VARCHAR(255),
//     creation_date DATE,
//     location VARCHAR(255),
//     responsible_person VARCHAR(255),
//     owner_name VARCHAR(255),
//     owner_address VARCHAR(255)
// );
