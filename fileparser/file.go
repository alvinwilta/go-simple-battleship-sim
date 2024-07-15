package fileparser

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

type Parser interface {
	Readfile(filepath string) string

	// WriteFile(filepath string)
}
