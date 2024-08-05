// graph/query/query.go
package query

import (
	"context"

	"pontistudios/gogogo/finance"
	"pontistudios/gogogo/graph/model"
)

type Resolver struct {
	Notes *finance.NotesResolver
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
