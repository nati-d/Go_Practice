package usecase

import (
	infrastructure "task_manager_testing/Infrastructure"
	"task_manager_testing/domain"

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

func (uu *UserUsecase) GetUserById(id primitive.ObjectID) (domain.User, error) {
	
	return uu.userRepo.GetUserById(id)
}

func (uu *UserUsecase) UpdateUser(id primitive.ObjectID, user domain.User) error {
	

	hashedPassword,err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	newUser := domain.User{ID: id, Username: user.Username, Password: string(hashedPassword), Role: user.Role}

	return uu.userRepo.UpdateUser(id, newUser)
}

func (uu *UserUsecase) DeleteUser(id primitive.ObjectID) error {
	return uu.userRepo.DeleteUser(id)
}

func (uu *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return uu.userRepo.GetAllUsers()
}
