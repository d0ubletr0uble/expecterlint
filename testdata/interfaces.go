package testdata

import "context"

type User struct {
	Name string
	Age  int
}

//go:generate go run github.com/vektra/mockery/v2@v2.52.2
type UserIFace interface {
	GetUser(ctx context.Context, name string) (User, error)
	CreateUser(context.Context, User) error
}
