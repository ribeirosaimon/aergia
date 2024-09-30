package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/logs"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia-utils/storage/pgsql"
	"github.com/ribeirosaimon/aergia/internal/config/database"
	"github.com/ribeirosaimon/aergia/internal/repository/mock"
)

type UserRepositoryInterface interface {
	InsertUser(ctx context.Context, user *sql.User) error
	FindUserByEmail(ctx context.Context, email string) (*sql.User, error)
}

var userOnce sync.Once
var userRepository UserRepositoryInterface
var userTable = "user"

// NewUserRepository is once open function
func NewUserRepository() UserRepositoryInterface {
	userOnce.Do(func() {
		switch properties.GetEnvironmentMode() {
		case constants.PROD, constants.DEV, constants.INTEGRATION:
			userRepository = newUserRepositoryImpl()
		default:
			userRepository = new(mock.UserRepositoryMock)
		}
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

func (u *UserRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*sql.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepositoryImpl) InsertUser(ctx context.Context, user *sql.User) error {
	// I have this func but I'm prioritizing performance
	query := createInsertQuery(user)
	logs.LOG.Message(query)
	exec, err := u.conn.GetConnection().Exec(query)
	if err != nil {
		logs.ERROR.Message(query)
		return err
	}
	exists, err := exec.RowsAffected()
	if err != nil {
		logs.ERROR.Message(query)
		return err
	}
	if exists == 0 {
		logs.ERROR.Message(query)
		return errors.New("cannot create user")
	}
	return nil
}

func createInsertQuery(user *sql.User) string {
	return fmt.Sprintf(`
	INSERT INTO "%s"
		(username, password, email, first_name, last_name, role)
	VALUES
		('%s', '%s', '%s', '%s', '%s', '%s')
	`, userTable,
		user.Username.GetValue(),
		user.Password.GetValue(),
		user.Email.GetValue(),
		user.FirstName.GetValue(),
		user.LastName.GetValue(),
		user.Role,
	)
}
