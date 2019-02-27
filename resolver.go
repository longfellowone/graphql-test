package graphql

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, todo TodoInput) (*Todo, error) {
	return &Todo{
		TodoID: "3423",
		Title:  todo.Title,
		Text:   todo.Text,
	}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	return []Todo{}, nil
}
func (r *queryResolver) Todo(ctx context.Context, todoID string) (*Todo, error) {
	panic("not implemented")
}
