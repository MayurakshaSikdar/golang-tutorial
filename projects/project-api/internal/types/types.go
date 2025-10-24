package types

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age"`
	Email string `json:"email" validate:"required"`
}
