package repository

import (
	"dakhou/todo/pkg/io"
	"github.com/jinzhu/gorm"
)

type todoRepo struct {
	DB *gorm.DB
}

func (t2 todoRepo) GetByID(id uint) (t []*io.Todo, err error) {
	panic("implement me")
}

func (t2 todoRepo) Get() (ts []*io.Todo, err error) {
	panic("implement me")
}

func (t2 todoRepo) Add(todo *io.Todo) (t *io.Todo, err error) {
	panic("implement me")
}

func (t2 todoRepo) Update(todo *io.Todo) (t *io.Todo, err error) {
	panic("implement me")
}

func (t2 todoRepo) Delete(id uint) (err error) {
	panic("implement me")
}

func (t2 todoRepo) ToggleComplete(id uint) (err error) {
	panic("implement me")
}

func (t2 todoRepo) SetStar(star uint8) (err error) {
	panic("implement me")
}

func (t2 todoRepo) FetchChildren(id uint) (ts []*io.Todo, err error) {
	panic("implement me")
}

func (t2 todoRepo) GetCategoryByID(id uint) (c *io.TodoCategory, err error) {
	panic("implement me")
}

func (t2 todoRepo) GetCategory() (cs []*io.TodoCategory, err error) {
	panic("implement me")
}

func (t2 todoRepo) AddCategory(category *io.TodoCategory) (c *io.TodoCategory, err error) {
	panic("implement me")
}

func (t2 todoRepo) UpdateCategory(category *io.TodoCategory) (c *io.TodoCategory, err error) {
	panic("implement me")
}

func (t2 todoRepo) DeleteCategory(id uint) (err error) {
	panic("implement me")
}

func (t2 todoRepo) FetchCategoryChildren(id uint) (cs []*io.TodoCategory, err error) {
	panic("implement me")
}

type Repository interface {
	//	todo methods:
	GetByID(id uint) (t []*io.Todo, err error)
	Get() (ts []*io.Todo, err error)
	Add(todo *io.Todo) (t *io.Todo, err error)
	Update(todo *io.Todo) (t *io.Todo, err error)
	Delete(id uint) (err error)
	ToggleComplete(id uint) (err error)
	SetStar(star uint8) (err error)
	FetchChildren(id uint) (ts []*io.Todo, err error)
	
	// category methods:
	GetCategoryByID(id uint) (c *io.TodoCategory, err error)
	GetCategory() (cs []*io.TodoCategory, err error)
	AddCategory(category *io.TodoCategory) (c *io.TodoCategory, err error)
	UpdateCategory(category *io.TodoCategory) (c *io.TodoCategory, err error)
	DeleteCategory(id uint) (err error)
	FetchCategoryChildren(id uint) (cs []*io.TodoCategory, err error)
}

func NewTodoRepo(db *gorm.DB) Repository {
	return &todoRepo{DB:db}
}
