package models

// Page contains details of the page
type Page struct {
	Number int
	Size   int
}

//StartIndex returns start index of a page
func (p Page) StartIndex() int {
	return (p.Number - 1) * p.Size
}

//EndPosition returns the end position of a page
func (p Page) EndPosition() int {
	return p.Number * p.Size
}