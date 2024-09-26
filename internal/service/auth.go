package service

import (
	"context"
	"regexp"
	"sync"

	"github.com/pkg/errors"
	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/role"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/domain/valueobject"
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
	dbUser.ID = valueobject.NewUuid()
	email, err := valueobject.NewEmail(user.Email)
	if err != nil {
		return err
	}
	dbUser.Email = *email
	dbUser.Username = valueobject.NewName(user.Username)
	password, err := valueobject.NewPassword(user.Password)
	if err != nil {
		return err
	}
	dbUser.Password = *password
	dbUser.LastName = valueobject.NewName(user.LastName)
	dbUser.FirstName = valueobject.NewName(user.FirstName)
	dbUser.Role = role.USER
	dbUser.Status = valueobject.PENDING
	dbUser.LoginAtempt = 0

	_, err = a.userRepository.CreateUser(ctx, &dbUser)
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
	}

	allowedUserPassword := 6
	if len(user.Password) < allowedUserPassword {
		return errors.New("password too short")
	}

	hasSpecialChar := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(user.Password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(user.Password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(user.Password)

	if !hasSpecialChar || !hasUppercase || !hasDigit {
		return errors.New("Password must contain uppercase and lower case, digit and special characters")
	}

	if user.Email == "" {
		return errors.New("email required")
	}

	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	if matchString := re.MatchString(user.Email); !matchString {
		return errors.New("email invalid")
	}

	return nil
}
