package entity

type Device struct {
	Id   string `json:"id"`
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}
