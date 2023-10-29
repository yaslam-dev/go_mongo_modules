// Package home provides middleware and shared auth reponsibilities
package home

// Auth function to authenticate user
func Auth(token string) bool {
	return true
}

// Authorize function to authorize user
func Authorize(token string) bool {
	return true
}
