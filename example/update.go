package example

import (
	"context"
	"fmt"
)

type UpdateBuilder struct {
	context    context.Context
	repository *Repository

	where []string
	args  []interface{}
	keys  []string
}

// TODO: Support expression
//  set({ age: () => "age + 1"  })
func (this *UpdateBuilder) Set(values map[string]interface{}) *UpdateBuilder {
	for key, value := range values {
		this.keys = append(this.keys, key)
		this.args = append(this.args, value)
	}

	return this
}

func (this *UpdateBuilder) Where(condition string, arguments ...interface{}) *UpdateBuilder {
	this.where = append(this.where, condition)
	this.args = append(this.args, arguments...)

	return this
}

func (this *UpdateBuilder) Execute() error {
	sql := "UPDATE users"

	sql += " SET"
	for i, key := range this.keys {
		if i != 0 {
			sql += ", "
		}

		sql += fmt.Sprintf(" %s = ?", key)
	}

	sql += " WHERE "
	for k, where := range this.where {
		if k != 0 {
			sql += " AND "
		}

		sql += fmt.Sprintf("(%s)", where)
	}

	fmt.Println("SQL: ", sql, this.args)

	return nil
}

func (this *UpdateBuilder) SQL() string {
	sql := "UPDATE users"

	sql += " SET"
	for i, key := range this.keys {
		if i != 0 {
			sql += ", "
		}

		sql += fmt.Sprintf(" %s = ?", key)
	}

	sql += " WHERE "
	for k, where := range this.where {
		if k != 0 {
			sql += " AND "
		}

		sql += fmt.Sprintf("(%s)", where)
	}

	return sql
}
