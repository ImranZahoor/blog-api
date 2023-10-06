package storage

import (
	"fmt"
	"sync"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	InMemoryStorage struct {
		articles map[models.Uuid]models.Article
		mutex    sync.Mutex
	}
)

var (
	ArticleAlreadyExists = fmt.Errorf("Article Alreay Exists")
)

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{articles: make(map[models.Uuid]models.Article)}
}

func (im *InMemoryStorage) Create(article models.Article) error {
	im.mutex.Lock()
	defer im.mutex.Unlock()
	id := im.nextKey()
	article.Id = id
	im.articles[id] = article
	return nil
}

func (im *InMemoryStorage) List() ([]models.Article, error) {
	var articles []models.Article
	for _, v := range im.articles {
		articles = append(articles, v)
	}
	return articles, nil
}

func (im *InMemoryStorage) nextKey() models.Uuid {

	return models.Uuid(len(im.articles) + 1)
}
