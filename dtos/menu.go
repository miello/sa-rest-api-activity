package dtos

type UpdateMenuBodyDtos struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type CreateMenuBodyDtos struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type CreateMenuResponseDtos struct {
	ID int `json:"id"`
}
