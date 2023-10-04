package repository

import (
	"context"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	ArticleRepository interface {
		GetArticleByID(ctx context.Context) (*models.Article, error)
		CreateArticle(ctx context.Context, article *models.Article) error
		ListArticle(ctx context.Context) ([]models.Article, error)
	}
)

func (r repository) GetArticleByID(ctx context.Context) (*models.Article, error) {
	return &models.Article{}, nil
}
func (r repository) CreateArticle(ctx context.Context, article *models.Article) error {
	return nil
}
func (r repository) ListArticle(ctx context.Context) ([]models.Article, error) {
	return []models.Article{}, nil
}
