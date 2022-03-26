package storage

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/post/model"
)

type storage struct {
	mysql.Repository
}

func New(repository mysql.Repository) Repository {
	return &storage{repository}
}

func (s *storage) Create(ctx context.Context, tx mysql.Transaction, post *model.Post) error {

	postDB := model.NewPostDB(post)

	insert := `
	INSERT INTO strider.posts(
			id,
			content,
			parent,
			created_at,
	        created_by
		)
	VALUES(?, ?, ?, ?, ?)`

	_, err := tx.ExecContext(
		ctx,
		insert,
		postDB.ID,
		postDB.Content,
		postDB.ParentID,
		postDB.CreatedAt,
		postDB.CreatedBy)

	return err
}

func (s *storage) List(ctx context.Context, limit, offset int) ([]*model.Post, error) {

	query := `
		SELECT 
			child.id,
			child.content,
			child.created_at,
			child.created_by,
		    parent.id,
			parent.content,
			parent.created_at,
			parent.created_by
		FROM
			strider.posts child
		LEFT JOIN strider.posts parent ON child.parent = parent.id
		ORDER BY child.created_at DESC
		LIMIT ?, ?`

	rows, err := s.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for rows.Next() {
		listDB := new(model.ListDBDTO)
		err = rows.Scan(
			&listDB.ID,
			&listDB.Content,
			&listDB.CreatedAt,
			&listDB.CreatedBy,
			&listDB.ParentID,
			&listDB.ParentContent,
			&listDB.ParentCreatedAt,
			&listDB.ParentCreatedBy)

		posts = append(posts, model.NewPost(listDB))
	}

	return posts, nil
}

func (s *storage) ListByFollower(ctx context.Context, follower string, limit, offset int) ([]*model.Post, error) {
	query := `
		SELECT 
			child.id,
			child.content,
			child.created_at,
			child.created_by,
		    parent.id,
			parent.content,
			parent.created_at,
			parent.created_by
		FROM
			strider.posts child
		LEFT JOIN strider.posts parent ON child.parent = parent.id
		WHERE child.created_by IN (SELECT following FROM strider.follower WHERE id = ?)
		ORDER BY child.created_at DESC
		LIMIT ?, ?`

	rows, err := s.QueryContext(ctx, query, follower, offset, limit)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for rows.Next() {
		listDB := new(model.ListDBDTO)
		err = rows.Scan(
			&listDB.ID,
			&listDB.Content,
			&listDB.CreatedAt,
			&listDB.CreatedBy,
			&listDB.ParentID,
			&listDB.ParentContent,
			&listDB.ParentCreatedAt,
			&listDB.ParentCreatedBy)

		posts = append(posts, model.NewPost(listDB))
	}

	return posts, nil
}

func (s *storage) ListByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Post, error) {
	query := `
		SELECT 
			child.id,
			child.content,
			child.created_at,
			child.created_by,
		    parent.id,
			parent.content,
			parent.created_at,
			parent.created_by
		FROM
			strider.posts child
		LEFT JOIN strider.posts parent ON child.parent = parent.id
		WHERE child.created_by = ?
		ORDER BY child.created_at DESC
		LIMIT ?, ?`

	rows, err := s.QueryContext(ctx, query, userID, offset, limit)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for rows.Next() {
		listDB := new(model.ListDBDTO)
		err = rows.Scan(
			&listDB.ID,
			&listDB.Content,
			&listDB.CreatedAt,
			&listDB.CreatedBy,
			&listDB.ParentID,
			&listDB.ParentContent,
			&listDB.ParentCreatedAt,
			&listDB.ParentCreatedBy)

		posts = append(posts, model.NewPost(listDB))
	}

	return posts, nil
}

func (s *storage) GetPostCountToday(ctx context.Context, userID string) (int, error) {
	query := `
		SELECT 
			count(1)
		FROM
			strider.posts
		WHERE created_by = ? and DATE_FORMAT(created_at,'%Y-%m-%d') = CURDATE()`

	row := s.QueryRowContext(ctx, query, userID)

	var postsCount int
	if err := row.Scan(&postsCount); err != nil {
		return 0, err
	}

	return postsCount, nil
}
