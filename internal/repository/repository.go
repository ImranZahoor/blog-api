package repository

import "github.com/ImranZahoor/blog-api/pkg/storage"

type (
	Repository interface {
		ArticleRepository
	}
	repository struct {
		memory *storage.InMemoryStorage
	}
)

func NewRepository(memStorage *storage.InMemoryStorage) Repository {
	return &repository{memory: memStorage}
}
