package models

type Uuid int

type Article struct {
	Id          Uuid   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Category struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
