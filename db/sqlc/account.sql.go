// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: account.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createAccount = `-- name: CreateAccount :one
insert into accounts (
    user_id, 
    category_id, 
    title,
    type,
    description,
    value,
    date
) values (
    $1, $2, $3, $4, $5, $6, $7
) returning id, user_id, category_id, title, type, description, value, date, created_at
`

type CreateAccountParams struct {
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Value       int32     `json:"value"`
	Date        time.Time `json:"date"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.UserID,
		arg.CategoryID,
		arg.Title,
		arg.Type,
		arg.Description,
		arg.Value,
		arg.Date,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
delete from accounts where id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
select id, user_id, category_id, title, type, description, value, date, created_at from accounts 
where id = $1 limit 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountGraph = `-- name: GetAccountGraph :one
select count(*) from accounts 
where user_id = $1 and type = $2
`

type GetAccountGraphParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountGraph(ctx context.Context, arg GetAccountGraphParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountGraph, arg.UserID, arg.Type)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAccountReports = `-- name: GetAccountReports :one
select sum(value) as sum_value from accounts 
where user_id = $1 and type = $2
`

type GetAccountReportsParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountReports(ctx context.Context, arg GetAccountReportsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountReports, arg.UserID, arg.Type)
	var sum_value int64
	err := row.Scan(&sum_value)
	return sum_value, err
}

const getAccounts = `-- name: GetAccounts :many
select 
    a.user_id, 
    a.category_id, 
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title as category_title
from accounts a
left join categories c on c.id = a.category_id 
where a.user_id = $1 and a.category_id = $2 and a.type = $3 
and a.title like $4 and a.description like $5 and a.date = $6
`

type GetAccountsParams struct {
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type GetAccountsRow struct {
	UserID        int32          `json:"user_id"`
	CategoryID    int32          `json:"category_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts,
		arg.UserID,
		arg.CategoryID,
		arg.Type,
		arg.Title,
		arg.Description,
		arg.Date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsRow{}
	for rows.Next() {
		var i GetAccountsRow
		if err := rows.Scan(
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
update accounts set title = $2, description = $3, value = $4 where id = $1 returning id, user_id, category_id, title, type, description, value, date, created_at
`

type UpdateAccountParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Value,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}
