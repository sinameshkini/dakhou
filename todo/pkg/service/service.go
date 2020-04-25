package service

import (
	"context"
	"dakhou/todo/pkg/io"
	"dakhou/todo/pkg/repository"
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

type basicTodoService struct{
	repo repository.Repository
}

func (b *basicTodoService) Get(ctx context.Context) (t []*io.Todo, error error) {
	// TODO implement the business logic of Get
	return t, error
}
func (b *basicTodoService) Add(ctx context.Context, todo *io.Todo) (t *io.Todo, error error) {
	// TODO implement the business logic of Add
	return t, error
}
func (b *basicTodoService) SetComplete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of SetComplete
	return error
}
func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of RemoveComplete
	return error
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of Delete
	return error
}
func (b *basicTodoService) Update(ctx context.Context, todo *io.Todo) (t *io.Todo, error error) {
	// TODO implement the business logic of Update
	return t, error
}
func (b *basicTodoService) SetStar(ctx context.Context, id string, star uint8) (error error) {
	// TODO implement the business logic of SetStar
	return error
}
func (b *basicTodoService) GetChildren(ctx context.Context, id string) (t []*io.Todo, error error) {
	// TODO implement the business logic of GetChildren
	return t, error
}
func (b *basicTodoService) GetCategory(ctx context.Context) (c []*io.TodoCategory, error error) {
	// TODO implement the business logic of GetCategory
	return c, error
}
func (b *basicTodoService) AddCategory(ctx context.Context, category *io.TodoCategory) (c *io.TodoCategory, error error) {
	// TODO implement the business logic of AddCategory
	return c, error
}
func (b *basicTodoService) UpdateCategory(ctx context.Context, category *io.TodoCategory) (c *io.TodoCategory, error error) {
	// TODO implement the business logic of UpdateCategory
	return c, error
}
func (b *basicTodoService) DeleteCategory(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of DeleteCategory
	return error
}
func (b *basicTodoService) GetCatChildes(ctx context.Context, id string) (c []*io.TodoCategory, error error) {
	// TODO implement the business logic of GetCatChildes
	return c, error
}

// NewBasicTodoService returns a naive, stateless implementation of TodoService.
func NewBasicTodoService(repository repository.Repository) TodoService {
	return &basicTodoService{
		repo: repository,
	}
}

// New returns a TodoService with all of the expected middleware wired in.
func New(repo repository.Repository, middleware []Middleware) TodoService {
	var svc TodoService = NewBasicTodoService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
