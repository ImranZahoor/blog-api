package models

type Article struct {
	Id          int64
	Title       string
	Description string
}

type Category struct {
	Id          int64
	Name        string
	Description string
}
