package service

type (
	Service interface {
		ArticleService
		CategoryService
	}

	service struct{}
)

func NewService() Service {
	return service{}
}
