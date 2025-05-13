package domain

import (
	"context"
)

type FileRepository interface {
	ReadFile(ctx context.Context, importJobFile ImportJobFile) ([]PlayerGameStatistic, error)
}
type ImportJobFileRepository interface {
	Get(ctx context.Context, id int) (*ImportJobFile, error)
	AddPlayerGameStatistics(ctx context.Context, playerGameStatistics []PlayerGameStatistic, jobID int) error
}

type InsertRawDataService struct {
	fileRepository          FileRepository
	importJobFileRepository ImportJobFileRepository
}

func NewInsertRawDataService(fileRepository FileRepository, importJobFileRepository ImportJobFileRepository) *InsertRawDataService {
	return &InsertRawDataService{fileRepository: fileRepository, importJobFileRepository: importJobFileRepository}
}

func (c *InsertRawDataService) Run(ctx context.Context, id int) error {

	importJobFile, err := c.importJobFileRepository.Get(ctx, id)
	if err != nil {
		return err
	}
	playerGameStatistics, err := c.fileRepository.ReadFile(ctx, *importJobFile)
	if err != nil {
		return err
	}
	err = c.importJobFileRepository.AddPlayerGameStatistics(ctx, playerGameStatistics, id)
	if err != nil {
		return err
	}

	return nil
}
