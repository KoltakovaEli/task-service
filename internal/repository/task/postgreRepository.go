package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"task-service/internal/entity"
	"time"
)

type postgreRepository struct {
	db *sql.DB
}

func createPointer(task entity.Task) *entity.Task {
	return &task
}
func NewPostgresRepository(db *sql.DB) *postgreRepository {
	return &postgreRepository{db}
}

func scanTasks(rows *sql.Rows) ([]*entity.Task, error) {
	var tasks []*entity.Task
	var t entity.Task
	for rows.Next() {
		err := rows.Scan(&t.ID, &t.Name, &t.UserID)
		tasks = append(tasks, createPointer(t))
		if err != nil {
			return []*entity.Task{}, fmt.Errorf("failed to scan tasks from db %w", err)
		}
	}
	return tasks, nil
}

func scanTask(row *sql.Row) (entity.Task, error) {
	var t entity.Task
	err := row.Scan(&t.ID, &t.Name, &t.UserID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("failed to scan tasks from db %w", err)
	}
	return t, nil
}

func (p postgreRepository) GetTasks(ctx context.Context) ([]*entity.Task, error) {
	sql := "SELECT id,name,user_id FROM tasks"
	rows, err := p.db.Query(sql)
	if err != nil {
		return []*entity.Task{}, err
	}
	tasks, err := scanTasks(rows)
	if err != nil {
		return []*entity.Task{}, err
	}
	return tasks, nil
}

func (p postgreRepository) Create(ctx context.Context, name, userID string) (*entity.Task, error) {
	if name == "" || userID == "" {
		return nil, errors.New("empty name or userID")
	}
	sql := "INSERT INTO tasks (name, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	params := []interface{}{
		name,
		userID,
		time.Now().UTC(),
		time.Now().UTC(),
	}
	row := p.db.QueryRow(sql, params...)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert task into db %w", err)
	}
	return &entity.Task{
		ID:     id,
		Name:   name,
		UserID: userID,
	}, nil
}

func (p postgreRepository) GetByID(ctx context.Context, ID string) (*entity.Task, error) {
	sql := "SELECT id,name,user_id FROM tasks WHERE id = $1"
	row := p.db.QueryRow(sql, ID)
	task, err := scanTask(row)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (p postgreRepository) Delete(ctx context.Context, ID string) error {
	sql := "DELETE FROM tasks WHERE id = $1;"
	_, err := p.db.Exec(sql, ID)
	if err != nil {
		return err
	}
	return nil
}

func (p postgreRepository) Update(ctx context.Context, ID string, Name string) (*entity.Task, error) {
	sql := "UPDATE tasks SET name = $1 WHERE id = $2 RETURNING id, user_id"
	row := p.db.QueryRow(sql, Name, ID)
	var id, userID string
	err := row.Scan(&id, &userID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert task into db %w", err)
	}
	return &entity.Task{
		ID:     id,
		Name:   Name,
		UserID: userID,
	}, nil
}
