package services

import (
	"AnalyseService/analytics"
	"fmt"
	"github.com/MaxAixs/protos/gen/api/gen/api"
	"github.com/google/uuid"
)

func convertToGRPCModel(tasks []analytics.CompletedTaskModel) (*api.WeeklyCompletedTasksResponse, error) {
	var grpcTasks []*api.CompletedTask

	for _, task := range tasks {
		grpcTasks = append(grpcTasks, &api.CompletedTask{
			UserId: task.UserId.String(),
			Email:  task.Email,
			Count:  task.Count,
		})
	}

	return &api.WeeklyCompletedTasksResponse{
		Tasks: grpcTasks,
	}, nil
}

func convertToTaskModel(task *api.TaskDoneItem) (analytics.TaskModel, error) {
	userID, err := uuid.Parse(task.UserId)
	if err != nil {
		return analytics.TaskModel{}, fmt.Errorf("cant parsing to UUID : %v", err)
	}

	return analytics.TaskModel{
		UserId:    userID,
		Email:     task.Email,
		ItemId:    int(task.ItemId),
		CreatedAt: task.CreatedAt.AsTime(),
	}, nil
}
