package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"gogogo/server/graph/model"
)

// ProcessIcs is the resolver for the processICS field.
func (r *mutationResolver) ProcessIcs(ctx context.Context, input *model.ProcessICSInput) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented: ProcessIcs - processICS"))
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// CalculateGoal is the resolver for the calculateGoal field.
func (r *mutationResolver) CalculateGoal(ctx context.Context, input model.BudgetInput) (*model.GoalCalculationOutput, error) {
	panic(fmt.Errorf("not implemented: CalculateGoal - calculateGoal"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Expenses is the resolver for the expenses field.
func (r *queryResolver) Expenses(ctx context.Context) ([]*model.Expense, error) {
	panic(fmt.Errorf("not implemented: Expenses - expenses"))
}

// Goals is the resolver for the goals field.
func (r *queryResolver) Goals(ctx context.Context) ([]*model.Goal, error) {
	panic(fmt.Errorf("not implemented: Goals - goals"))
}

// Budget is the resolver for the budget field.
func (r *queryResolver) Budget(ctx context.Context) (*model.Budget, error) {
	panic(fmt.Errorf("not implemented: Budget - budget"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
