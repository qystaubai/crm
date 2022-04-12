package users

import (
	crm "github.com/qystaubai/crm/entities"
)

var (
	ErrorAccessDenied string = "ErrorAccessDenied"
	ErrorNotFound     string = "ErrorNotFound"
	ErrorForbidden    string = "ErrorForbidden"
)

type UsersCore interface {
	ListUsers() (*[]crm.User, error)
}

type UsersCoreImpl struct {
	repo crm.UsersStorager
}

func NewUsersCoreImpl(repo crm.UsersStorager) *UsersCoreImpl {
	return &UsersCoreImpl{
		repo,
	}
}

func (pa *UsersCoreImpl) ListUsers() (*[]crm.User, error) {
	// TODO IMPLEMENT
	return pa.repo.ListAll()
}
