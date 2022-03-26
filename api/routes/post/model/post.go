package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Request struct {
	Content   string `json:"content" validate:"lte=777"`
	CreatedBy string `json:"created_by" validate:"required"`
	ParentID  string `json:"parent"`
}

type Post struct {
	ID        string `json:"id,omitempty"`
	Content   string `json:"content,omitempty"`
	Parent    *Post  `json:"parent,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func NewPost(listDBDTO *ListDBDTO) *Post {
	post := new(Post)
	post.ID = listDBDTO.ID.String
	if listDBDTO.Content.Valid {
		post.Content = listDBDTO.Content.String
	}
	if listDBDTO.CreatedAt.Valid {
		post.CreatedAt = listDBDTO.CreatedAt.Time.Format(time.RFC3339)
	}
	if listDBDTO.CreatedBy.Valid {
		post.CreatedBy = listDBDTO.CreatedBy.String
	}

	parent := new(Post)
	if listDBDTO.ParentID.Valid {
		parent.ID = listDBDTO.ParentID.String
	}
	if listDBDTO.ParentContent.Valid {
		parent.Content = listDBDTO.ParentContent.String
	}
	if listDBDTO.ParentCreatedAt.Valid {
		parent.CreatedAt = listDBDTO.ParentCreatedAt.Time.Format(time.RFC3339)
	}
	if listDBDTO.ParentCreatedBy.Valid {
		parent.CreatedBy = listDBDTO.ParentCreatedBy.String
	}

	post.Parent = parent
	return post
}

type PostDB struct {
	ID        sql.NullString
	Content   sql.NullString
	ParentID  sql.NullString
	CreatedBy sql.NullString
	CreatedAt sql.NullTime
}

func NewPostDB(post *Post) *PostDB {

	postDB := new(PostDB)
	post.ID = uuid.New().String()
	postDB.ID.String = post.ID
	postDB.ID.Valid = true
	postDB.CreatedBy.String = post.CreatedBy
	postDB.CreatedBy.Valid = true
	postDB.CreatedAt.Time = time.Now()
	postDB.CreatedAt.Valid = true
	post.CreatedAt = postDB.CreatedAt.Time.Format(time.RFC3339)

	if len(post.Content) > 0 {
		postDB.Content.String = post.Content
		postDB.Content.Valid = true
	}

	if post.Parent != nil {
		postDB.ParentID.String = post.Parent.ID
		postDB.ParentID.Valid = true
	}

	return postDB
}

type ListRequest struct {
	User     string `query:"user"`
	Follower string `query:"follower"`
	Offset   int    `query:"offset"`
	Limit    int    `query:"limit"`
}

type ListDBDTO struct {
	ID              sql.NullString
	Content         sql.NullString
	CreatedBy       sql.NullString
	CreatedAt       sql.NullTime
	ParentID        sql.NullString
	ParentContent   sql.NullString
	ParentCreatedBy sql.NullString
	ParentCreatedAt sql.NullTime
}
