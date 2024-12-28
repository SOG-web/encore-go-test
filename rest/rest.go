package rest

import (
	"context"
)

type REST struct {
	Message string
}

// Get retrieves rest word.
//
//encore:api public method=GET path=/rest
func Get(ctx context.Context) (*REST, error) {

	return &REST{Message: "rest"}, nil
}
