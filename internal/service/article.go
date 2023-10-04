package service

import "context"

type (
	ArticleService interface {
		ListArticle(ctx context.Context)
		GetArticleByID(ctx context.Context)
		CreateArticle(ctx context.Context)
		UpdateArticle(ctx context.Context)
		DeleteArticle(ctx context.Context)
	}
)

func (s service) ListArticle(ctx context.Context)    {}
func (s service) GetArticleByID(ctx context.Context) {}
func (s service) DeleteArticle(ctx context.Context)  {}
func (s service) CreateArticle(ctx context.Context)  {}
func (s service) UpdateArticle(ctx context.Context)  {}
