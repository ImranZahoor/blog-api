package repository

import (
	"context"
	"fmt"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	ArticleRepository interface {
		GetArticleByID(ctx context.Context) (*models.Article, error)
		CreateArticle(ctx context.Context, article models.Article) error
		ListArticle(ctx context.Context) ([]models.Article, error)
	}
)

func (r repository) GetArticleByID(ctx context.Context) (*models.Article, error) {
	return &models.Article{}, nil
}
func (r repository) CreateArticle(ctx context.Context, article models.Article) error {
	err := r.memory.Create(article)
	if err != nil {
		return fmt.Errorf("failed to create articel %v", err)
	}
	return nil
}
func (r repository) ListArticle(ctx context.Context) ([]models.Article, error) {
	articles, err := r.memory.List()
	if err != nil {
		return []models.Article{}, err
	}
	return articles, nil
}
