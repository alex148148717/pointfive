package import_worker

import (
	"context"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/fx"
	"pointfive/internal/config"
	"pointfive/internal/import_worker/domain"
	"pointfive/internal/import_worker/infrastructure"
	"pointfive/internal/import_worker/interfaces"
)

var ImportWorkerModule = fx.Options(
	fx.Provide(
		interfaces.NewImportWorker,
		interfaces.NewParseDataActivity,
		infrastructure.NewLocalFileRepository,
		infrastructure.NewPlayerGameStatisticRepository,
		domain.NewInsertRawDataService,
	),
	fx.Invoke(func(lc fx.Lifecycle, config *config.Config,
		workerClient client.Client,
		importWorker *interfaces.ImportWorker,
		insertRawDataActivity *interfaces.InsertRawDataActivity) {
		w := worker.New(workerClient, config.TaskQueueImportJob, worker.Options{})

		w.RegisterWorkflowWithOptions(importWorker.ImportWorkerFlow, workflow.RegisterOptions{
			Name: "importJob",
		})

		w.RegisterActivityWithOptions(insertRawDataActivity.ParseDataActivity, activity.RegisterOptions{
			Name: "ParseDataActivity",
		})

		//w.RegisterActivity(activities.HelloActivity)
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := w.Run(worker.InterruptCh()); err != nil {
						panic(err)
					}
				}()
				return nil
			},
		})

	}),
)
