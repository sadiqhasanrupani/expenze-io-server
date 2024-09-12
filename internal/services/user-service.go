package services

import "expenze-io.com/internal/repositories"

type UserService struct {
  repo repositories.UserRepository
}

func NewUserService() *UserService {
  return &UserService{
    repo: *repositories.NewUserRepository(),
  }
}


