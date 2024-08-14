package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/ribeirosaimon/aergia-utils/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/logs"
	"github.com/ribeirosaimon/aergia-utils/storage/pgsql"
	"github.com/ribeirosaimon/aergia/internal/config/database"
)

var userOnce sync.Once
var userRepository UserRepositoryInterface
var userTable = "user"

// NewUserRepository is once open function
func NewUserRepository() UserRepositoryInterface {
	userOnce.Do(func() {
		userRepository = newUserRepositoryImpl()
	})
	return userRepository
}

func newUserRepositoryImpl() UserRepositoryInterface {
	return &UserRepositoryImpl{
		conn: database.NewPgsqlConnection(),
	}
}

type UserRepositoryImpl struct {
	conn pgsql.AergiaPgsqlInterface
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, user *sql.User) (*sql.User, error) {
	// I have this func but I'm prioritizing performance
	// query := u.conn.CreateQuery(user)
	query := createInsertQuery(user)
	logs.LOG.Message(query)
	exec, err := u.conn.GetConnection().Exec(query, userTable)
	if err != nil {
		logs.ERROR.Message(query)
		return nil, err
	}
	exists, err := exec.RowsAffected()
	if err != nil {
		logs.ERROR.Message(query)
		return nil, err
	}
	if exists == 0 {
		logs.ERROR.Message(query)
		return nil, errors.New("cannot create user")
	}
	return user, nil
}

func createInsertQuery(user *sql.User) string {
	return fmt.Sprintf(`
	INSERT INTO "%s"
		(username, password, email, first_name, last_name, role)
	VALUES
		('%s', '%s', '%s', '%s', '%s', '%s')
	`, userTable,
		user.Username,
		user.Password,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Role)
}

func (u *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*sql.User, error) {
	// TODO implement me
	panic("implement me")
}
