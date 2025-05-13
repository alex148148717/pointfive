package infrastructure

import (
	"context"
	"fmt"
	"io"
	"os"
	"pointfive/internal/config"
	"pointfive/internal/import_job/domain"
)

type LocalFileRepository struct {
}

func NewLocalFileRepository(config *config.Config) domain.FileRepository {
	r := LocalFileRepository{}
	return &r
}

func (c *LocalFileRepository) UploadFile(ctx context.Context) (io.WriteCloser, string, error) {
	//add uuid +machine name and etc
	tmpFile, err := os.CreateTemp("", "raw_file-*.txt")
	if err != nil {
		return nil, "", fmt.Errorf("failed to create temp file: %w", err)
	}
	return tmpFile, tmpFile.Name(), nil
}
