package models

// DefaultSuccess godoc
type DefaultSuccess struct {
	Message []string `json:"messages"`
}

// DefaultError godoc
type DefaultError struct {
	Error []string `json:"errors"`
}

// AuthError godoc
type AuthError struct {
	Error []string `json:"authErrors"`
}
