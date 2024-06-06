package models

type Course struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `json:"name"`
	Length      float32 `json:"length"`
	Keywords    string  `json:"keywords"`
	Description string  `json:"description"`
	Req         string  `json:"req"`
	Teachername string  `json:"teachername"`
}
