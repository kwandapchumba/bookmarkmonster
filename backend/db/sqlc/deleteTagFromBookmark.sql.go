// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: deleteTagFromBookmark.sql

package db

import (
	"context"
)

const deleteTagFromBookmark = `-- name: DeleteTagFromBookmark :exec
delete from bookmark_tag where bookmark_id = $1 and tag_id = $2
`

type DeleteTagFromBookmarkParams struct {
	BookmarkID string `json:"bookmark_id"`
	TagID      string `json:"tag_id"`
}

func (q *Queries) DeleteTagFromBookmark(ctx context.Context, arg DeleteTagFromBookmarkParams) error {
	_, err := q.db.Exec(ctx, deleteTagFromBookmark, arg.BookmarkID, arg.TagID)
	return err
}