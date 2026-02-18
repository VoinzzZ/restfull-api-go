package usecase

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"restfull-api-go/internal/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	existing, err := u.userRepo.FindByEmail(user.Email)
	if err != nil {
		return err
	}

	if existing != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user.Password = string(hashedPassword)

	if err != nil {
		return errors.New("failed to hash password")
	}

	return u.userRepo.Create(user)
}

func (u *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	return u.userRepo.FindAll()
}

func (u *UserUsecase) UpdateUser(id int, updateData *domain.User) error {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	user.Name = updateData.Name
	user.Email = updateData.Email

	return u.userRepo.Update(user)
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.userRepo.Delete(id)
}
