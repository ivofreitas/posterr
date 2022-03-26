package model

import (
	"database/sql"
	"time"
)

type Follower struct {
	ID        string `json:"id,omitempty"`
	Follow    string `json:"follow,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type FollowerDB struct {
	ID        sql.NullString
	Follow    sql.NullString
	CreatedAt sql.NullTime
}

func NewFollowerDB(follower *Follower) *FollowerDB {
	followerDB := new(FollowerDB)
	followerDB.ID.String = follower.ID
	followerDB.ID.Valid = true
	followerDB.Follow.String = follower.Follow
	followerDB.Follow.Valid = true
	followerDB.CreatedAt.Time = time.Now()
	followerDB.CreatedAt.Valid = true
	follower.CreatedAt = followerDB.CreatedAt.Time.Format(time.RFC3339)

	return followerDB
}

type Request struct {
	Follower string `json:"follower" query:"follower" validate:"required"`
	Follow   string `json:"follow" query:"follow" validate:"required"`
}
