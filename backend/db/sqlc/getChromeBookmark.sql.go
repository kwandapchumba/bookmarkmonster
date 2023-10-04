// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: getChromeBookmark.sql

package db

import (
	"context"
)

const getChromeBookmark = `-- name: GetChromeBookmark :one
select id, title, bookmark, host, favicon, thumbnail, notes, user_id, added, updated, deleted, folder_id, beautified, fromchrome from bookmark where bookmark = $1 and title = $2 and user_id = $3 and fromChrome = 'true' and deleted is null limit 1
`

type GetChromeBookmarkParams struct {
	Bookmark string `json:"bookmark"`
	Title    string `json:"title"`
	UserID   string `json:"user_id"`
}

func (q *Queries) GetChromeBookmark(ctx context.Context, arg GetChromeBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, getChromeBookmark, arg.Bookmark, arg.Title, arg.UserID)
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
		&i.FolderID,
		&i.Beautified,
		&i.Fromchrome,
	)
	return i, err
}
