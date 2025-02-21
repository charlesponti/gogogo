// graph/query/query.go
package query

import (
	"context"

	notes "gogogo/notes"
	"gogogo/server/graph/model"
)

type Resolver struct {
	Notes *notes.NotesResolver
}

func (r *Resolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{
			ID:   "1",
			Name: "Alice",
		},
		{
			ID:   "2",
			Name: "Bob",
		},
	}, nil
}
