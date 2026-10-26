package repository

import (
	"context"
	"time"

	"todo_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(parentCtx context.Context, pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer cancel()

	var query string = "INSERT INTO todos (title, completed, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id, title, completed, created_at, updated_at"
	var todo models.Todo
	err := pool.QueryRow(ctx, query, title, completed).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}