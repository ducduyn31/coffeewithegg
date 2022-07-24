package domain

import "context"

type ProjectDTO struct {
	ID int64 `json:"id"`
}

type ProjectRepository interface {
	Fetch(ctx context.Context) (res []ProjectDTO, err error)
}
