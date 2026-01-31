package service

import (
	"time"

	"github.com/gabrielssssssssss/marketplace-telegram/helper"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *accountServiceImpl) RegisterUser(user *entity.Users) (*model.Users, error) {
	users := entity.Users{
		UserId:      user.UserId,
		Username:    user.Username,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Balance:     0.0,
		RecoveryKey: helper.RandomStringSecure(24),
		UpdatedAt:   time.Now(),
	}

	resp, err := s.repository.CreateUser(&users)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *accountServiceImpl) FindUser(UserID int64) (*model.Users, error) {
	users := entity.Users{
		UserId: UserID,
	}

	resp, err := s.repository.GetUserByID(&users)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *accountServiceImpl) UpdateUserBalance(user *entity.Users) (*model.Users, error) {
	users := entity.Users{
		UserId:      user.UserId,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Username:    user.Username,
		Balance:     user.Balance,
		RecoveryKey: user.RecoveryKey,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	resp, err := s.repository.UpdateUserByID(&users)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
