// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: getRootFolderByName.sql

package db

import (
	"context"
)

const getRootFolderByName = `-- name: GetRootFolderByName :one
select exists(select folder_id, user_id, folder_name, path, label, starred, created_at, updated_at, subfolder_of, deleted_at, folder_description, folder_password, ishidden from folder where folder_name = $1 and subfolder_of is null limit 1)
`

func (q *Queries) GetRootFolderByName(ctx context.Context, folderName string) (bool, error) {
	row := q.db.QueryRow(ctx, getRootFolderByName, folderName)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
