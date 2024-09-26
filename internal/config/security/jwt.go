package security

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia-utils/util"
)

func CreateToken(loggedUser *util.LoggerUser) (string, error) {

	loggedByte, err := json.Marshal(loggedUser)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": loggedByte,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(properties.GetEnvironmentValue(constants.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(ctx context.Context, tokenString string) (context.Context, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return properties.GetEnvironmentValue(constants.SecretKey), nil
	})

	if err != nil {
		return ctx, err
	}

	if !token.Valid {
		return ctx, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userBytes, ok := claims["user"].(string); ok {
			var loggedUser util.LoggerUser
			if err = json.Unmarshal([]byte(userBytes), &loggedUser); err != nil {
				return ctx, err
			}
			return context.WithValue(ctx, constants.LoggedUser, loggedUser), nil
		}
	}
	return ctx, errors.New("invalid user")
}
