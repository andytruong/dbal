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

func (this *Repository) Select(ctx context.Context) *SelectBuilder {
	return &SelectBuilder{
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

func (this *Repository) Update(ctx context.Context) *UpdateBuilder {
	return &UpdateBuilder{
		context:    ctx,
		repository: this,
		where:      []string{},
		args:       []interface{}{},
		keys:       []string{},
	}
}
