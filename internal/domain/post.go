package domain

import "time"

type Post struct {
	Id        string
	Slug      string
	Title     string
	AuthorId  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FindPostRequest struct {
}

type FindPostResponse struct {
}

type CreatePostRequest struct {
}

type CreatePostResponse struct {
}

type DeletePostRequest struct {
}

type DeletePostResponse struct {
}

type UpdatePostRequest struct {
}

type UpdatePostResponse struct {
}
