package usecase

import (
	infrastructure "task_manager_refactored/Infrastructure"
	"task_manager_refactored/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uu *UserUsecase) RegisterUser(username, password, role string) error {
	hashedPassword,err := infrastructure.HashPassword(password)
	if err != nil {
		return err
	}

	
	return uu.userRepo.RegisterUser(username, string(hashedPassword), role)
}

func (uu *UserUsecase) Login(username, password string) (domain.User, error) {
	return uu.userRepo.Login(username, password)
}

func (uu *UserUsecase) GetUserById(id string) (domain.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, err
	}
	return uu.userRepo.GetUserById(oid)
}

func (uu *UserUsecase) UpdateUser(id, username, password, role string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	hashedPassword,err := infrastructure.HashPassword(password)
	if err != nil {
		return err
	}
	user := domain.User{ID: oid, Username: username, Password: string(hashedPassword), Role: role}

	return uu.userRepo.UpdateUser(oid, user)
}

func (uu *UserUsecase) DeleteUser(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return uu.userRepo.DeleteUser(oid)
}

func (uu *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return uu.userRepo.GetAllUsers()
}
