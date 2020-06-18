package example

import (
	"context"
	"database/sql"
)

type (
	Repository struct {
		db *sql.DB
	}
)

func (this *Repository) QueryBuilder(ctx context.Context) *QueryBuilder {
	return &QueryBuilder{
		context:    ctx,
		repository: this,
	}
}
