package handler

import (
	mockhandler "AnalyseService/analytics/handler/mocks"
	"errors"
	"github.com/MaxAixs/protos/gen/api/gen/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestHandler_SaveDoneTasks(t *testing.T) {
	type mockBehavior func(m *mockhandler.MockAnalytics, ctx context.Context, req *api.TaskDoneItems)

	tests := []struct {
		name          string
		input         *api.TaskDoneItems
		mockBehavior  mockBehavior
		expectMessage string
		expectError   bool
	}{
		{
			name: "success",
			input: &api.TaskDoneItems{
				Items: []*api.TaskDoneItem{
					{
						UserId:    "Test123",
						ItemId:    123,
						Email:     "test@test.com",
						CreatedAt: timestamppb.Now(),
					},
				},
			},
			mockBehavior: func(m *mockhandler.MockAnalytics, ctx context.Context, req *api.TaskDoneItems) {
				m.EXPECT().SaveTaskData(ctx, req).Return(nil)
			},
			expectMessage: "Analytic items data saved successfully",
			expectError:   false,
		}, {
			name: "failure - SaveTaskData returns error",
			input: &api.TaskDoneItems{
				Items: []*api.TaskDoneItem{
					{
						UserId:    "Test123",
						ItemId:    123,
						Email:     "test@test.com",
						CreatedAt: timestamppb.Now(),
					},
				},
			},
			mockBehavior: func(m *mockhandler.MockAnalytics, ctx context.Context, req *api.TaskDoneItems) {
				m.EXPECT().SaveTaskData(ctx, req).Return(errors.New("failed to save data"))
			},
			expectMessage: "",
			expectError:   true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			mockAnalytics := mockhandler.NewMockAnalytics(ctrl)

			tt.mockBehavior(mockAnalytics, ctx, tt.input)

			apiHandler := &AnalyticsAPI{
				Analytics: mockAnalytics,
			}

			resp, err := apiHandler.SaveDoneTasks(ctx, tt.input)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectMessage, resp.Message)
			}
		})
	}
}

func TestHandler_FetchWeeklyCompletedTask(t *testing.T) {
	type mockBehavior func(m *mockhandler.MockAnalytics, ctx context.Context, req *emptypb.Empty)

	tests := []struct {
		name           string
		input          *emptypb.Empty
		mockBehavior   mockBehavior
		expectResponse *api.WeeklyCompletedTasksResponse
		expectError    bool
	}{
		{
			name:  "success",
			input: &emptypb.Empty{},
			mockBehavior: func(m *mockhandler.MockAnalytics, ctx context.Context, req *emptypb.Empty) {
				m.EXPECT().GetWeeklyList(ctx).Return(&api.WeeklyCompletedTasksResponse{Tasks: []*api.CompletedTask{}}, nil)
			},
			expectResponse: &api.WeeklyCompletedTasksResponse{Tasks: []*api.CompletedTask{}},
			expectError:    false,
		},
		{
			name:  "failure - GetWeeklyList returns error",
			input: &emptypb.Empty{},
			mockBehavior: func(m *mockhandler.MockAnalytics, ctx context.Context, req *emptypb.Empty) {
				m.EXPECT().GetWeeklyList(ctx).Return(nil, errors.New("failed to get weekly list"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			mockAnalytics := mockhandler.NewMockAnalytics(ctrl)

			tt.mockBehavior(mockAnalytics, ctx, tt.input)

			apiHandler := &AnalyticsAPI{
				Analytics: mockAnalytics,
			}

			resp, err := apiHandler.FetchWeeklyCompletedTask(ctx, tt.input)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectResponse, resp)
			}
		})
	}
}
