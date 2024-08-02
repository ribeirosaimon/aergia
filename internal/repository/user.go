package repository

import "sync"

var userOnce sync.Once
var userRepository UserRepositoryInterface

// NewUserRepository is once open function
func NewUserRepository() UserRepositoryInterface {
	userOnce.Do(func() {
		userRepository = newUserrepositoryImpl()
	})
	return userRepository
}

func newUserrepositoryImpl() UserRepositoryInterface {
}

type UserRepositoryImpl struct {
}
