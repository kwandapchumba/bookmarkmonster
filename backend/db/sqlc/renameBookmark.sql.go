// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: renameBookmark.sql

package db

import (
	"context"
)

const renameBookmark = `-- name: RenameBookmark :one
update bookmark set title = $1 where id = $2 returning id, title, bookmark, host, favicon, thumbnail, notes, user_id, added, updated, deleted
`

type RenameBookmarkParams struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}

func (q *Queries) RenameBookmark(ctx context.Context, arg RenameBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, renameBookmark, arg.Title, arg.ID)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Bookmark,
		&i.Host,
		&i.Favicon,
		&i.Thumbnail,
		&i.Notes,
		&i.UserID,
		&i.Added,
		&i.Updated,
		&i.Deleted,
	)
	return i, err
}
