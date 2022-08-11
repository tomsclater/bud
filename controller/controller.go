package controller

import (
	context "context"
)

// Controller for stories
type Controller struct {
}

// Story struct
type Story struct {
	ID int `json:"id"`
}

// Index of stories
// GET
func (c *Controller) Index(ctx context.Context) (stories []*Story, err error) {
	return []*Story{}, nil
}

// Show story
// GET :id
func (c *Controller) Show(ctx context.Context, id int) (story *Story, err error) {
	return &Story{
		ID: id,
	}, nil
}
