package infrastructure

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"pointfive/internal/config"
	"pointfive/internal/import_job/domain"
)

type WorkerRepository struct {
	workerClient       client.Client
	taskQueueImportJob string
}

func NewWorkerRepository(workerClient client.Client, config *config.Config) domain.ImportWorkerRepository {
	r := WorkerRepository{
		workerClient:       workerClient,
		taskQueueImportJob: config.TaskQueueImportJob,
	}
	return &r
}

func (c *WorkerRepository) InsertWorkerJob(ctx context.Context, id int) (*domain.WorkerJob, error) {
	key := fmt.Sprintf("importJob_%d", id)
	options := client.StartWorkflowOptions{
		ID:        key,
		TaskQueue: c.taskQueueImportJob,
	}
	we, err := c.workerClient.ExecuteWorkflow(ctx, options, "importJob", id)
	if err != nil {
		return nil, err
	}
	WorkerJob := &domain.WorkerJob{ID: we.GetID(), RunID: we.GetRunID()}
	return WorkerJob, nil
}
