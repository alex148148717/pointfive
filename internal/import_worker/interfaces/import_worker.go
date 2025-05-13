package interfaces

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/workflow"
	"log"
	"os"
	"os/exec"
	"time"
)

type ImportWorker struct {
	insertRawDataActivity *InsertRawDataActivity
}

func NewImportWorker(insertRawDataActivity *InsertRawDataActivity) *ImportWorker {
	h := ImportWorker{insertRawDataActivity: insertRawDataActivity}
	return &h
}

func (c *ImportWorker) ImportWorkerFlow(ctx workflow.Context, id int) error {
	logger := workflow.GetLogger(ctx)
	logger.Info(fmt.Sprintf("XXX Import worker flow started id %d\n", id))

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, "ParseDataActivity", id).Get(ctx, &result)
	if err != nil {
		return err
	}

	logger.Info("Workflow finished", "result", result)

	return nil
}
func (c *ImportWorker) ParseDataActivity(ctx context.Context, id int) (string, error) {
	msg := fmt.Sprintf("ParseDataActivity, %d!", id)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)

	}
	fmt.Println("Current working directory:", dir)
	cmd := exec.Command("/Users/alexgreenman/.local/bin/dbt", "run",
		"--project-dir", dir+"/dbt",
		"--profiles-dir", dir+"/dbt/profiles",
		"--models", "my_dbt_project.alex",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//cmd.Env = append(os.Environ(), "DBT_PROFILES_DIR=/path/to/profiles")

	if err := cmd.Run(); err != nil {
		log.Fatalf("dbt run failed: %v", err)
	}
	return msg, nil
}
