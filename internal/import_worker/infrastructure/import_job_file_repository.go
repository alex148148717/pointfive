package infrastructure

import (
	"context"
	"database/sql"
	"pointfive/internal/config"
	"pointfive/internal/import_job/domain"
	"time"
)

type ImportJobFileRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB, config *config.Config) domain.ImportJobFileRepository {
	r := ImportJobFileRepository{
		db: db,
	}
	return &r
}

func (c *ImportJobFileRepository) InsertImportJobFile(ctx context.Context, path string) (*domain.ImportJobFile, error) {
	query := `INSERT INTO import_job_files (path, time) VALUES ($1, $2) RETURNING id,path`
	var importJobFile domain.ImportJobFile
	err := c.db.QueryRow(query, path, time.Now()).Scan(&importJobFile.JobID, &importJobFile.Path)
	return &importJobFile, err
}
