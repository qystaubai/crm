package crm

type User struct {
	Name string
}

type UsersStorager interface {
	ListAll() (*[]User, error)
}
