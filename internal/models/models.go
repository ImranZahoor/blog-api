package models

type Uuid int

type Article struct {
	Id          Uuid   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Category struct {
	Id          Uuid   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	Id    Uuid   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
