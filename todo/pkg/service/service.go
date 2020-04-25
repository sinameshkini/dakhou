package service

import (
	"context"
	"dakhou/todo/pkg/io"
)

// TodoService describes the service.
type TodoService interface {
	// todo methods
	Get(ctx context.Context) (t []*io.Todo, error error)
	Add(ctx context.Context, todo *io.Todo) (t *io.Todo, error error)
	SetComplete(ctx context.Context, id string) (error error)
	RemoveComplete(ctx context.Context, id string) (error error)
	Delete(ctx context.Context, id string) (error error)
	Update(ctx context.Context, todo *io.Todo) (t *io.Todo, error error)
	SetStar(ctx context.Context, id string, star uint8) (error error)
	GetChildren(ctx context.Context, id string) (t []*io.Todo, error error)
	// Category methods
	GetCategory(ctx context.Context) (c []*io.TodoCategory, error error)
	AddCategory(ctx context.Context, category *io.TodoCategory) (c *io.TodoCategory, error error)
	UpdateCategory(ctx context.Context, category *io.TodoCategory) (c *io.TodoCategory, error error)
	DeleteCategory(ctx context.Context, id string) (error error)
	GetCatChildes(ctx context.Context, id string) (c []*io.TodoCategory, error error)
}
