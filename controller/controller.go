package controller

import (
	context "context"

	"github.com/matthewmueller/hackernews"
)

// Controller for stories
type Controller struct {
	HN *hackernews.Client
}

// // Generated Story Struct
// **REMOVED THIS SNIPPET**
// type Story struct {
// 	ID int `json:"id"`
// }

// Index of stories
// GET
func (c *Controller) Index(ctx context.Context) (stories []*hackernews.Story, err error) {
	return c.HN.FrontPage(ctx)
	// return []*Story{}, nil
}

// Show story
// GET :id
func (c *Controller) Show(ctx context.Context, id int) (story *hackernews.Story, err error) {
	return c.HN.Find(ctx, id)
	// **REMOVED THIS SNIPPET**
	// return &Story{
	// 	ID: id,
	// }, nil
}
