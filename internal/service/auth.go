package service

import (
	"context"
	"regexp"
	"sync"

	"github.com/pkg/errors"
	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/entities/role"
	"github.com/ribeirosaimon/aergia-utils/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/ribeirosaimon/aergia/internal/repository"
	"github.com/ribeirosaimon/aergia/internal/service/mock"
)

var authOnce sync.Once
var authService AuthServiceInterface

// NewAuthService is once open function
func NewAuthService() AuthServiceInterface {
	authOnce.Do(func() {
		switch properties.GetEnvironmentMode() {
		case constants.PROD, constants.DEV:
			authService = newAuthServiceImpl()
		default:
			authService = new(mock.AuthServiceMock)
		}

	})

	return authService
}

type authServiceImpl struct {
	userRepository repository.UserRepositoryInterface
}

func newAuthServiceImpl() AuthServiceInterface {
	return &authServiceImpl{
		userRepository: repository.NewUserRepository(),
	}
}

func (a *authServiceImpl) SignUp(ctx context.Context, user *dto.User) error {
	if err := a.userValidator(user); err != nil {
		return err
	}

	var dbUser sql.User

	dbUser.Email = user.Email
	dbUser.Username = user.Username
	dbUser.Password = user.Password

	dbUser.Role = role.USER
	dbUser.LastName = user.LastName
	dbUser.FirstName = user.FirstName
	dbUser.LoginAtempt = 0

	_, err := a.userRepository.CreateUser(ctx, &dbUser)
	if err != nil {
		return err
	}
	return nil
}

func (a *authServiceImpl) Login(ctx context.Context, login, pass string) error {
	return nil
}

func (a *authServiceImpl) userValidator(user *dto.User) error {
	if user.Password == "" {
		return errors.New("password required")
	} else {
		hasSpecialChar := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(user.Password)
		hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(user.Password)
		hasDigit := regexp.MustCompile(`[0-9]`).MatchString(user.Password)

		if !hasSpecialChar || !hasUppercase || !hasDigit {
			return errors.New("Password must contain uppercase and lower case characters")
		}

	}
	if user.Email == "" {
		return errors.New("email required")
	} else {
		regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(regex)
		re.MatchString(user.Email)
	}
	return nil
}
