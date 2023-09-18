package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"mado/internal/core/survey"
	"mado/pkg/database/postgres"
	"mado/pkg/logger"
)

// Survey is a Survey repository.
type SurveyrRepository struct {
	db *postgres.Postgres
}

// NewSurveyRepository creates a new UserRepository.
func NewSurveyrRepository(db *postgres.Postgres) SurveyrRepository {
	return SurveyrRepository{
		db: db,
	}
}

// todo implement this
func (s SurveyrRepository) Create(req *survey.SurveyRequirements, ctx context.Context) (*survey.SurveyRequirements, error) {

	// Create a transaction to ensure atomicity
	tx, err := s.db.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted, // Set the isolation level to Serializable
		AccessMode:     pgx.ReadWrite,     // Set the access mode to ReadWrite
		DeferrableMode: pgx.NotDeferrable, // Set the deferrable mode to NotDeferrable
	})
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}
	//defer statement is used to ensure that a rollback is initiated if any error occurs during the transaction
	defer func() {
		if err != nil {
			// An error occurred, rollback the transaction
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				// Handle rollback error if needed
				fmt.Printf("Error rolling back transaction: %v\n", rollbackErr)
			}
		}
	}()

	// Create a slice to store the generated question IDs
	questionIDs := []int{}

	// Insert questions into the 'public.question' table and retrieve the generated IDs
	for _, q := range req.Questions {
		// Insert a question and retrieve the generated ID
		var questionID int
		err := tx.QueryRow(ctx, "INSERT INTO public.question (description) VALUES ($1) RETURNING id", q.Description).Scan(&questionID)
		if err != nil {
			return nil, fmt.Errorf("could not insert question: %w", err)
		}
		questionIDs = append(questionIDs, questionID)
	}

	// todo use real fields
	mockRka := "Mock Rka Value"
	mockRcName := "Mock RcName Value"
	mockAdress := "Mock Address Value"

	// Use the Squirrel query builder to create the SQL query
	// Assuming you have a 'survey' table with appropriate columns
	insertBuilder := s.db.Builder.Insert("public.survey").
		Columns("name", "rka", "rc_name", "adress", "question_id", "user_id").
		Values(req.Name, mockRka, mockRcName, mockAdress, questionIDs, req.UserID).
		Suffix("RETURNING id")

	// Build the SQL query
	sqlQuery, args, err := insertBuilder.ToSql()
	if err != nil {
		// fmt.Printf("logger.NewLevel(\"error\"): %v\n", logger.NewLevel("error"))
		return &survey.SurveyRequirements{}, fmt.Errorf("can not build insert survey query: %w", err)
	}

	logger.FromContext(ctx).Debug("check following query", zap.String("sql", sqlQuery), zap.Any("args", args))

	// Execute the query and retrieve the generated ID
	var id string
	if err = s.db.Pool.QueryRow(ctx, sqlQuery, args...).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return &survey.SurveyRequirements{}, fmt.Errorf("can not insert user: %w", survey.ErrAlreadyExist)
			}
		}
		return &survey.SurveyRequirements{}, fmt.Errorf("can not insert survey: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	req.ID = id
	return req, nil

}
