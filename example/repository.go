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

func (this *Repository) Create(ctx context.Context, obj *User) error {
	_, err := this.db.ExecContext(
		ctx,
		"INSERT INTO users (id, name) VALUES (?, ?)",
		obj.ID, obj.Name,
	)

	return err
}
