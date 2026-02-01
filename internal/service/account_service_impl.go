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
		Role:        user.Role,
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

func (s *accountServiceImpl) FindUserByID(userID int64) (*model.Users, error) {
	users := entity.Users{
		UserId: userID,
	}

	resp, err := s.repository.GetUserByID(&users)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *accountServiceImpl) FindUserByRecoveryKey(recoveryKey string) (*model.Users, error) {
	users := entity.Users{
		RecoveryKey: recoveryKey,
	}

	resp, err := s.repository.GetUserByRecoveryKey(&users)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *accountServiceImpl) ModifyUserByID(user *entity.Users) (*model.Users, error) {
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

func (s *accountServiceImpl) RemoveUserByID(UserID int64) (bool, error) {
	users := entity.Users{
		UserId: UserID,
	}

	_, err := s.repository.DeleteUser(&users)
	if err != nil {
		return false, err
	}

	return true, nil
}
