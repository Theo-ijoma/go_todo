package repository

import (
	"context"
	"todo_api/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, completed bool)(models.Todo, error) {
	// Implement the logic to insert the todo into the database
	// For example, using a SQL INSERT statement with the provided context and todo data
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	return models.Todo{}, nil
}