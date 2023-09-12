// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: getEmailConfirmationLink.sql

package db

import (
	"context"
)

const getEmailConfirmationLink = `-- name: GetEmailConfirmationLink :one
select email, code, expiry, user_password from email_verification where email = $1
`

func (q *Queries) GetEmailConfirmationLink(ctx context.Context, email string) (EmailVerification, error) {
	row := q.db.QueryRow(ctx, getEmailConfirmationLink, email)
	var i EmailVerification
	err := row.Scan(
		&i.Email,
		&i.Code,
		&i.Expiry,
		&i.UserPassword,
	)
	return i, err
}