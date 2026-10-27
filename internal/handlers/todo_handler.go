package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"todo_api/internal/repository"
)

type CreateTodoRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func CreateTodoInput(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateTodoRequest

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := repository.CreateTodo(c.Request.Context(), pool, input.Title, input.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
			return
		}

		c.JSON(http.StatusCreated, todo)

	}
}

func GetTodos(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := repository.GetTodos(c.Request.Context(), pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
			return
		}
		c.JSON(http.StatusOK, todos)
	}
}
