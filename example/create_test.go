package example

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Mock_Repository() Repository {
	db, err := sql.Open("sqlite3", ":memory:")

	if nil != err {
		panic(err)
	}

	_, err = db.Exec(
		`
			CREATE TABLE users (
    			id   uuid                  NOT NULL PRIMARY KEY,
    			name character varying(64) NOT NULL
            )`,
	)

	if nil != err {
		panic(err)
	}

	return Repository{db: db}
}

func Test_Create(t *testing.T) {
	ass := assert.New(t)
	ctx := context.Background()
	r := Mock_Repository()
	defer r.db.Close()

	obj := &User{
		ID:   "john.doe",
		Name: "John Doe",
	}

	err := r.Create(ctx, obj)
	ass.NoError(err)

	{
		obj, err := r.Select(ctx).
			Where("id = ?", "john.doe").
			GetOne()

		ass.NoError(err)
		ass.Equal(obj.ID, "john.doe")
		ass.Equal(obj.Name, "John Doe")
	}
}
