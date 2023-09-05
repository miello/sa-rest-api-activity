package dtos

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseError struct {
	Message string `json:"message"`
}
