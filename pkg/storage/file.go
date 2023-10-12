package storage

import (
	"bytes"
	"encoding/gob"
	"os"

	"github.com/ImranZahoor/blog-api/internal/models"
)

type (
	FileStorage struct {
		fileName   string
		fileHndler *os.File
	}
)

var (
	categories map[models.Uuid]models.Category
)

func NewFileStorage(fileName string) (*FileStorage, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	_ = make(map[models.Uuid]models.Category)
	fileStorage := &FileStorage{fileHndler: file, fileName: fileName}
	return fileStorage, nil
}

func (f *FileStorage) Create(category models.Category) error {

	id := models.Uuid(len(categories))
	category.Id = id
	categories[id] = category

	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)
	if err := e.Encode(categories); err != nil {
		return err
	}
	_, err := buf.WriteTo(f.fileHndler)
	if err != nil {
		return err
	}
	return nil
}
func (f *FileStorage) List() ([]models.Category, error) {

	var buf bytes.Buffer
	var categories map[models.Uuid]models.Category
	_, err := buf.ReadFrom(f.fileHndler)
	if err != nil {
		return []models.Category{}, err
	}
	e := gob.NewDecoder(&buf)
	if err := e.Decode(&categories); err != nil {
		return []models.Category{}, err
	}

	categoriesList := make([]models.Category, len(categories)-1)
	for _, v := range categories {
		categoriesList = append(categoriesList, v)
	}
	return categoriesList, nil
}

func (f *FileStorage) GetByID() error {
	return nil
}

func (f *FileStorage) Update() error {
	return nil
}

func (f *FileStorage) Delete() error {
	return nil
}
