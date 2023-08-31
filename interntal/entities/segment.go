package entities

type Segment struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}
