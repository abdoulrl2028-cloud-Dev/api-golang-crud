package service

import (
	"github.com/abdoulrl2028-cloud-Dev/api-golang-crud/internal/model"
	"github.com/abdoulrl2028-cloud-Dev/api-golang-crud/internal/repository"
)

// UserService handles user business logic
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.repo.GetByID(id)
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(id int, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	err = s.repo.Update(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
