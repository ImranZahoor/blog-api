package service

import (
	"context"
	"fmt"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	ArticleService interface {
		ListArticle(ctx context.Context) ([]models.Article, error)
		GetArticleByID(ctx context.Context)
		CreateArticle(ctx context.Context, article models.Article) error
		UpdateArticle(ctx context.Context)
		DeleteArticle(ctx context.Context)
	}
)

func (s service) ListArticle(ctx context.Context) ([]models.Article, error) {
	articles, err := s.repository.ListArticle(ctx)
	if err != nil {
		return []models.Article{}, err
	}
	return articles, nil
}
func (s service) GetArticleByID(ctx context.Context) {}
func (s service) DeleteArticle(ctx context.Context)  {}
func (s service) CreateArticle(ctx context.Context, article models.Article) error {
	err := s.repository.CreateArticle(ctx, article)
	if err != nil {
		return fmt.Errorf("error creating article %s", err)
	}
	return nil
}
func (s service) UpdateArticle(ctx context.Context) {}
