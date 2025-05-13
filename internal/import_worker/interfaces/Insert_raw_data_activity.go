package interfaces

import (
	"context"
	"pointfive/internal/import_worker/domain"
)

type InsertRawDataActivity struct {
	insertRawDataService *domain.InsertRawDataService
}

func NewParseDataActivity(insertRawDataService *domain.InsertRawDataService) *InsertRawDataActivity {
	h := InsertRawDataActivity{insertRawDataService: insertRawDataService}
	return &h
}

func (c *InsertRawDataActivity) ParseDataActivity(ctx context.Context, id int) (string, error) {
	c.insertRawDataService.Run(ctx, id)

	return "s", nil
}
