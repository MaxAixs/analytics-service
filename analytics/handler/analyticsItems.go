package handler

import (
	"context"
	"github.com/MaxAixs/protos/gen/api/gen/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Analytics interface {
	SaveTaskData(ctx context.Context, items *api.TaskDoneItems) error
	GetWeeklyList(ctx context.Context) (*api.WeeklyCompletedTasksResponse, error)
}

type AnalyticsAPI struct {
	api.UnimplementedAnalyticsDataServer
	Analytics Analytics
}

func Register(gRPC *grpc.Server, analytics Analytics) {
	api.RegisterAnalyticsDataServer(gRPC, &AnalyticsAPI{Analytics: analytics})

}

func (a *AnalyticsAPI) SaveDoneTasks(ctx context.Context, req *api.TaskDoneItems) (*api.ServiceResponse, error) {
	err := a.Analytics.SaveTaskData(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.ServiceResponse{Message: "Analytic items data saved successfully"}, nil
}

func (a *AnalyticsAPI) FetchWeeklyCompletedTask(ctx context.Context, req *emptypb.Empty) (*api.WeeklyCompletedTasksResponse, error) {
	tasks, err := a.Analytics.GetWeeklyList(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return tasks, nil
}
