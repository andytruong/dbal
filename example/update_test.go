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

	// Setup data for update.
	{
		obj := &User{
			ID:   "john.doe",
			Name: "John Doe",
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
			Where("id = ?", "john.doe").
			SQL()

		ass.Equal("UPDATE users SET name = ? WHERE (id = ?)", sql)
	})
}
