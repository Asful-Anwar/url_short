package repository

import (
	"database/sql"
	"github.com/Asful-Anwar/url-shortener/model"
)

type LinkRepository struct {
	DB *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{DB: db}
}

func (r *LinkRepository) CreateLink(link *model.Link) error {
	query := `INSERT INTO links (link, newlink, user_id, title, description, created_at, modified_at, expired_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW())`
	_, err := r.DB.Exec(query, link.Link, link.Newlink, link.UserID, link.Title, link.Description)
	return err
}