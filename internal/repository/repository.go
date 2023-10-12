package repository

import "github.com/ImranZahoor/blog-api/pkg/storage"

type (
	Repository interface {
		ArticleRepository
		UserRepository
	}
	repository struct {
		memory *storage.InMemoryStorage
		db     *storage.MySQLStorage
	}
)

func NewRepository(memStorage *storage.InMemoryStorage, dbStorage *storage.MySQLStorage) Repository {
	return &repository{memory: memStorage, db: dbStorage}

}
