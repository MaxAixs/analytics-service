package repository

import (
	"AnalyseService/analytics"
	"context"
	"database/sql"
	"fmt"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (t *TaskRepo) SaveTask(ctx context.Context, modelTask analytics.TaskModel) error {
	q := `INSERT INTO TaskDoneItem (UserID, ItemID, Email, CreatedAt) VALUES ($1, $2, $3, $4)`

	_, err := t.db.ExecContext(ctx, q, modelTask.UserId, modelTask.ItemId, modelTask.Email, modelTask.CreatedAt)
	if err != nil {
		return fmt.Errorf("can save Task: %w", err)
	}

	return nil
}

func (t *TaskRepo) LoadTask(ctx context.Context) ([]analytics.CompletedTaskModel, error) {
	var tasks []analytics.CompletedTaskModel

	querySelect := `
		SELECT UserID, Email, COUNT(*) AS TotalCompleted
		FROM TaskDoneItem
		WHERE CreatedAt >= NOW() - INTERVAL '7 days'
		  AND Sent_notify = false
		GROUP BY UserID, Email;
	`

	rows, err := t.db.QueryContext(ctx, querySelect)
	if err != nil {
		return nil, fmt.Errorf("db err : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task analytics.CompletedTaskModel
		if err := rows.Scan(&task.UserId, &task.Email, &task.Count); err != nil {
			return nil, fmt.Errorf("scan err: %w", err)
		}
		tasks = append(tasks, task)
	}

	if len(tasks) > 0 {
		if err := t.updateSentNotify(ctx); err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func (t *TaskRepo) updateSentNotify(ctx context.Context) error {
	queryUpdate := `
		UPDATE TaskDoneItem
		SET Sent_notify = true
		WHERE CreatedAt >= NOW() - INTERVAL '7 days'
		  AND Sent_notify = false;
	`

	_, err := t.db.ExecContext(ctx, queryUpdate)
	if err != nil {
		return fmt.Errorf("failed to update TaskDoneItem: %v", err)
	}

	return nil
}
