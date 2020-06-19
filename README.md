DBAL generator
====

[WIP] Generator for Go database access layer

## Why code generator?

Current tools like GORM using a lot of reflection to resolve the query, it's nice but it's much slower than raw query. It also:

- A lot of magic inside.
- Hard to debug
- Depends on a lot more libraries

## Usage

```golang
db, err := sql.Open(driver, dsn)
re := Repository{db: db}

// Create
err := re.Create(User { Id: id, Name: name })

// Update
err := re.Update(ctx).
    Where("id = ?", id).
    Set(map[string]interface{}{ "name": "Jon" }).
    Execute()

// Load
obj, err := r.Select(ctx).Where("id = ?", id).GetOne()
```

## Status

>>> WIP
