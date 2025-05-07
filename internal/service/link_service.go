package service

import (
	"math/rand"
	"time"
	"github.com/Asful-Anwar/url-shortener/model"
	"github.com/Asful-Anwar/url-shortener/internal/repository"
)

type LinkService struct {
	Repo *repository.LinkRepository
}

func NewLinkService(repo *repository.LinkRepository) *LinkService {
	return &LinkService{Repo: repo}
}

func (s *LinkService) GenerateShortLink() string {
	const charset ="loremipsumdolorsitametconsecteturadipiscingelit"
	b := make([]byte, 6)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *LinkService) CreateShortLink(originalLink string) (string, error){
	short := s.GenerateShortLink()
	link := model.Link{
		Link: originalLink,
		Newlink: short,
		UserID: 0,
		Title: "",
		Description: "",
	}
	err := s.Repo.CreateLink(&link)
	if err != nil {
		return "", err
	}
	return short, nil
}