package github

type Snippet struct {
}

type Client interface {
	Search() []Snippet
}

func NewClient() {
}
