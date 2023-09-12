// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: updateFolder.sql

package db

import (
	"context"
)

const updateFolder = `-- name: UpdateFolder :one
update folder set folder_name = $1, folder_description = $2 where folder_id = $3 and user_id = $4 returning folder_id, user_id, folder_name, path, label, starred, created_at, updated_at, subfolder_of, deleted_at, folder_description, folder_password, ishidden
`

type UpdateFolderParams struct {
	FolderName        string `json:"folder_name"`
	FolderDescription string `json:"folder_description"`
	FolderID          string `json:"folder_id"`
	UserID            string `json:"user_id"`
}

func (q *Queries) UpdateFolder(ctx context.Context, arg UpdateFolderParams) (Folder, error) {
	row := q.db.QueryRow(ctx, updateFolder,
		arg.FolderName,
		arg.FolderDescription,
		arg.FolderID,
		arg.UserID,
	)
	var i Folder
	err := row.Scan(
		&i.FolderID,
		&i.UserID,
		&i.FolderName,
		&i.Path,
		&i.Label,
		&i.Starred,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SubfolderOf,
		&i.DeletedAt,
		&i.FolderDescription,
		&i.FolderPassword,
		&i.Ishidden,
	)
	return i, err
}