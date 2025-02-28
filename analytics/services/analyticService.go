package services

import (
	"AnalyseService/analytics"
	"context"
	"fmt"
	"github.com/MaxAixs/protos/gen/api/gen/api"
	"sync"
)

type AnalyticService struct {
	itemsSaver  ItemsSaver
	itemsLoader ItemsLoader
}

type ItemsSaver interface {
	SaveTask(ctx context.Context, item analytics.TaskModel) error
}

type ItemsLoader interface {
	LoadTask(ctx context.Context) ([]analytics.CompletedTaskModel, error)
}

func NewAnalyticService(save ItemsSaver, load ItemsLoader) *AnalyticService {
	return &AnalyticService{itemsSaver: save, itemsLoader: load}
}

func (a *AnalyticService) SaveTaskData(ctx context.Context, items *api.TaskDoneItems) []error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(items.Items))

	for _, task := range items.Items {
		wg.Add(1)
		go func(task *api.TaskDoneItem) {
			defer wg.Done()

			if err := a.processSaveTaskData(ctx, task); err != nil {
				errChan <- err
			}
		}(task)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		errs := getErrFromChan(errChan)
		return errs
	}

	return nil
}

func getErrFromChan(errChan chan error) []error {
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	return errs
}

func (a *AnalyticService) processSaveTaskData(ctx context.Context, task *api.TaskDoneItem) error {
	taskModel, err := convertToTaskModel(task)
	if err != nil {
		return fmt.Errorf("cant convet task model %v", err)
	}

	return a.itemsSaver.SaveTask(ctx, taskModel)
}

func (a *AnalyticService) GetWeeklyList(ctx context.Context) (*api.WeeklyCompletedTasksResponse, error) {
	tasks, err := a.itemsLoader.LoadTask(ctx)
	if err != nil {
		return nil, fmt.Errorf("cant load tasks %v", err)
	}

	grpcTasks, err := convertToGRPCModel(tasks)
	if err != nil {
		return nil, fmt.Errorf("cant convert to grpc model %v", err)
	}

	return grpcTasks, nil
}
