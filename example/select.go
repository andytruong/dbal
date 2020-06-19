package example

import (
	"context"
	"fmt"
	"strconv"
)

type SelectBuilder struct {
	context    context.Context
	repository *Repository

	where []string
	order string
	args  []interface{}
	first int
}

func (this *SelectBuilder) First(first int) *SelectBuilder {
	this.first = first

	return this
}

func (this *SelectBuilder) Order(order string) *SelectBuilder {
	this.order = order

	return this
}

func (this *SelectBuilder) Where(condition string, arguments ...interface{}) *SelectBuilder {
	this.where = append(this.where, condition)
	this.args = append(this.args, arguments...)

	return this
}

func (this *SelectBuilder) GetOne() (*User, error) {
	sql := this.First(1).SQL()

	rows, err := this.repository.db.QueryContext(this.context, sql, this.args...)
	if nil != err {
		return nil, err
	}

	obj := &User{}
	for rows.Next() {
		err := rows.Scan(&obj.ID, &obj.Name)
		if err != nil {
			return nil, err
		}
	}

	return obj, err
}

func (this *SelectBuilder) GetMany() ([]User, error) {
	sql := this.First(1).SQL()

	rows, err := this.repository.db.QueryContext(this.context, sql, this.args...)
	if nil != err {
		return nil, err
	}

	list := []User{}
	for rows.Next() {
		obj := User{}
		err := rows.Scan(&obj.ID, &obj.Name)
		if err != nil {
			return nil, err
		}

		list = append(list, obj)
	}

	return list, err
}

func (this *SelectBuilder) SQL() string {
	sql := "SELECT id, name FROM users WHERE "

	if nil != this.where {
		for i, where := range this.where {
			if i > 0 {
				sql += " AND "
			}

			sql += fmt.Sprintf("(%s)", where)
		}
	}

	if this.order != "" {
		sql += " ORDER BY " + this.order
	}

	if this.first != 0 {
		sql += " LIMIT " + strconv.Itoa(this.first)
	}

	return sql
}
