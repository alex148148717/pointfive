package infrastructure

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"pointfive/internal/config"
	"pointfive/internal/import_worker/domain"
)

type LocalFileRepository struct {
}

func NewLocalFileRepository(config *config.Config) domain.FileRepository {
	r := LocalFileRepository{}
	return &r
}

func (c *LocalFileRepository) ReadFile(ctx context.Context, importJobFile domain.ImportJobFile) ([]domain.PlayerGameStatistic, error) {

	p := importJobFile.Path
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var entities []domain.PlayerGameStatistic
	if err := json.Unmarshal(bytes, &entities); err != nil {
		return nil, err
	}
	return entities, nil
}
