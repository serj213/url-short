package repo

import "errors"


var(
	ErrNotFound = errors.New("not found")
	ErrAliasExists = errors.New("alias exists")
)

const(
	PgCodeDublicate = "23505"
)