package entity

type Category struct {
	ID        uint
	Name      string
	Questions []Question
}
