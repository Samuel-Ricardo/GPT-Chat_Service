package entity

type Model struct {
	Name      string
	MaxTokens int
}

func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}
}


