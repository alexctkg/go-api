package models

// DefaultSuccess godoc
type DefaultSuccess struct {
	Message []string `json:"messages"`
}

// DefaultError godoc
type DefaultError struct {
	Error []string `json:"errors"`
}
