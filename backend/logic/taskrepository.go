package logic

import (
	"database/sql"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Count() (int, error)
	CreateTask(description string) error
}

type postgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) TaskRepository {
	return &postgresTaskRepository{db}
}

func (ptr *postgresTaskRepository) Count() (res int, err error) {
	row := ptr.db.QueryRow("SELECT COUNT(1) FROM tasks")

	res = -1
	err = row.Scan(&res)
	return
}

func (ptr *postgresTaskRepository) CreateTask(description string) error {
	_, err := ptr.db.Exec(
		"INSERT INTO tasks (task_id, task_description) VALUES ($1, $2)",
		uuid.NewString(),
		description,
	)
	return err
}
