package pq

import (
	// "github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
	// "github.com/pkg/errors"
	crm "github.com/qystaubai/crm/entities"
)

// Errors constants
const (
	ErrorSQLNoRowsAffected = "ErrorSQLNoRowsAffected"
	ErrorSQLNoRowsFound    = "ErrorSQLNoRowsFound"
)

// const tableName = "users"

// UsersStore represents
type UsersStore struct {
	// db *sqlx.DB
}

// New constructor
func New(
// db *sqlx.DB
) *UsersStore {
	return &UsersStore{
		// db,
	}
}

func (us *UsersStore) ListAll() (*[]crm.User, error) {
	var users []crm.User

	user := &crm.User{}

	user.Name = "TEST1"

	users = append(users, *user)

	// TODO: IMPLEMENT
	// err := us.db.QueryRowx(`SELECT * FROM users`).Scan(
	// 	...
	// )
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, nil
	// 	}
	// 	return nil, errors.WithStack(err)
	// }

	// for _, id := range []int64(tagIds) {
	// 	person.TagIDs = append(person.TagIDs, int(id))
	// }

	return &users, nil
}
