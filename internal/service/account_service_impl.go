package service

import (
	"github.com/gabrielssssssssss/marketplace-telegram/internal/entity"
	"github.com/gabrielssssssssss/marketplace-telegram/internal/model"
)

func (s *accountServiceImpl) RegisterUser(user *entity.User) (*model.User, error) {
	// users := entity.Users{
	// 	UserId:      user.UserId,
	// 	Username:    user.Username,
	// 	Firstname:   user.Firstname,
	// 	Lastname:    user.Lastname,
	// 	Role:        user.Role,
	// 	Balance:     0.0,
	// 	RecoveryKey: helper.RandomStringSecure(24),
	// 	UpdatedAt:   time.Now(),
	// }
	return s.repository.CreateUser(user)
}

func (s *accountServiceImpl) FindUserByID(user *entity.User) (*model.User, error) {
	return s.repository.GetUserByID(user)
}

func (s *accountServiceImpl) FindUserByRecoveryKey(user *entity.User) (*model.User, error) {
	return s.repository.GetUserByRecoveryKey(user)
}

func (s *accountServiceImpl) ModifyUserByID(user *entity.User) (*model.User, error) {
	return s.repository.UpdateUserByID(user)
}

func (s *accountServiceImpl) RemoveUserByID(user *entity.User) (bool, error) {
	return s.repository.DeleteUser(user)
}
