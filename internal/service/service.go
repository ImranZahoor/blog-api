package service

import "github.com/ImranZahoor/blog-api/internal/repository"

type (
	Service interface {
		ArticleService
		CategoryService
	}

	service struct {
		repository repository.Repository
	}
)

func NewService(repo repository.Repository) Service {
	return service{
		repository: repo,
	}
}
