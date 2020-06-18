package example

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ass := assert.New(t)
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "db.sqlite3")
	ass.NoError(err)
	id := "xxxxxxx"
	name := "John Do"

	t.Run("Where()", func(t *testing.T) {
		b := QueryBuilder{}
		query := b.
			Where("id = ? AND name = ?", id, name).
			SQL()
		ass.Equal("SELECT id, name FROM users WHERE (id = ? AND name = ?)", query)
		ass.Equal(id, b.args[0])
		ass.Equal(name, b.args[1])

	})

	t.Run("mutliple Where()", func(t *testing.T) {
		b := QueryBuilder{}

		query := b.
			Where("id = ?", id).
			Where("name = ?", name).
			SQL()

		ass.Equal("SELECT id, name FROM users WHERE (id = ?) AND (name = ?)", query)
		ass.Equal(id, b.args[0])
		ass.Equal(name, b.args[1])
	})

	t.Run("multiple Where() with different position", func(t *testing.T) {
		b := QueryBuilder{}

		query := b.
			Where("name = ?", name).
			Where("id = ?", id).
			SQL()

		ass.Equal("SELECT id, name FROM users WHERE (name = ?) AND (id = ?)", query)
		ass.Equal(name, b.args[0])
		ass.Equal(id, b.args[1])
	})

	t.Run("::Order()", func(t *testing.T) {
		b := QueryBuilder{}

		query := b.
			Where("name = ?", name).
			Where("id = ?", id).
			Order("id DESC").
			SQL()

		ass.Equal("SELECT id, name FROM users WHERE (name = ?) AND (id = ?) ORDER BY id DESC", query)
		ass.Equal(name, b.args[0])
		ass.Equal(id, b.args[1])
	})

	t.Run("::First()", func(t *testing.T) {
		b := QueryBuilder{}

		query := b.
			Where("name = ?", name).
			Where("id = ?", id).
			First(1).
			SQL()

		ass.Equal("SELECT id, name FROM users WHERE (name = ?) AND (id = ?) LIMIT 1", query)
		ass.Equal(name, b.args[0])
		ass.Equal(id, b.args[1])
	})

	t.Run("GetOne()", func(t *testing.T) {
		r := Repository{db: db}

		obj, err := r.QueryBuilder(ctx).Where("id = ?", id).
			Where("name = ?", name).GetOne()
		ass.NoError(err)
		ass.Equal(id, obj.ID)
		ass.Equal(name, obj.Name)
	})

	t.Run("GetMany()", func(t *testing.T) {
		r := Repository{db: db}

		list, err := r.QueryBuilder(ctx).Where("id = ?", id).
			Where("name = ?", name).
			Order("id DESC").
			GetMany()

		ass.NoError(err)
		ass.Equal(id, list[0].ID)
		ass.Equal(name, list[0].Name)
	})
}
