package repository

type (
	Repository interface {
		ArticleRepository
	}
	repository struct {
	}
)

func NewRepository() Repository {
	return &repository{}
}
