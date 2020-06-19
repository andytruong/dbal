package example

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Update(t *testing.T) {
	ass := assert.New(t)
	ctx := context.Background()
	r := Mock_Repository()
	defer r.db.Close()

	id := "john.doe"
	name := "John Doe"

	// Setup data for update.
	{
		obj := &User{
			ID:   id,
			Name: name,
		}

		err := r.Create(ctx, obj)
		ass.NoError(err)
	}

	t.Run("Check SQL query", func(t *testing.T) {
		sql := r.
			Update(ctx).
			Set(map[string]interface{}{
				"name": "Jon Do",
			}).
			Where("id = ?", id).
			SQL()

		ass.Equal("UPDATE users SET name = ? WHERE (id = ?)", sql)
	})

	t.Run("Check result", func(t *testing.T) {
		// update user
		err := r.Update(ctx).
			Where("id = ?", id).
			Set(map[string]interface{}{
				"name": "Jon Do",
			}).
			Execute()

		ass.NoError(err)

		// reload and check value
		obj, err := r.Select(ctx).Where("id = ?", id).GetOne()
		ass.NoError(err)

		ass.Equal(id, obj.ID)
		ass.Equal("Jon Do", obj.Name)
	})
}
