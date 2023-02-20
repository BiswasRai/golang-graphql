package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/biswasRai/golang-graphql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var isDone bool = false

	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Title:  input.Title,
		IsDone: &isDone,
	}

	if err := r.DB.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.TodoInput, id string) (*model.Todo, error) {
	todo := &model.Todo{
		ID:     id,
		Title:  input.Title,
		IsDone: input.IsDone,
	}

	if err := r.DB.Where("id = ?", id).Save(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	var todo *model.Todo

	if err := r.DB.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo

	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

// GetTodo is the resolver for the getTodo field.
func (r *queryResolver) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	var todo *model.Todo

	if err := r.DB.Where("id = ?", id).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }