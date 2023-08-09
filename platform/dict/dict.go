package dict

type Word struct {
	Word       string   `json:"word"`
	Definition []string `json:"definition"`
	Type       []string `json:"type"`
}

type Dict struct {
	words []Word
}

func NewDict() *Dict {
	return &Dict{
		words: []Word{},
	}
}
