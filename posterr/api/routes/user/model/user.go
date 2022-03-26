package model

import "database/sql"

type UpdateRequest struct {
	UserID         string
	PostsCount     int
	FollowerCount  int
	FollowingCount int
}

type UserDB struct {
	ID             sql.NullString
	Username       sql.NullString
	FollowerCount  sql.NullInt64
	FollowingCount sql.NullInt64
	PostsCount     sql.NullInt64
	CreatedAt      sql.NullTime
}
