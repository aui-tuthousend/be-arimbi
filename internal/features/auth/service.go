package auth

import (
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"
	"be-arimbi/utils"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(req *LoginRequest) (LoginResponse, error)
	RegisterUser(req user.UserRequest) error
}

type AuthServiceImpl struct {
	ar AuthRepository
	ur user.UserRepository
	rr role.RoleRepository
}

func NewAuthService(ar AuthRepository, ur user.UserRepository, rr role.RoleRepository) AuthService {
	return &AuthServiceImpl{ar: ar, ur: ur, rr: rr}
}

func (as *AuthServiceImpl) Login(req *LoginRequest) (LoginResponse, error) {
	existingUser, err := as.ar.FindUserByEmail(&req.Email)
	if err != nil {
		return LoginResponse{}, err
	}
	if existingUser == nil {
		return LoginResponse{}, errors.New("user not found")
	}
	if !as.ar.CheckPasswordHash(&req.Password, &existingUser.Password) {
		return LoginResponse{}, errors.New("wrong password")
	}

	existingRole, err := as.rr.GetByUuid(existingUser.RoleUuid)
	if err != nil {
		return LoginResponse{}, err
	}

	token, err := as.ar.GenerateJWT(&existingUser.Uuid, &existingRole.Name)
	if err != nil {
		return LoginResponse{}, err
	}

	userResponse := user.UserLoginResponse{
		Uuid: existingUser.Uuid,
		Name: existingUser.Name,
		Email: existingUser.Email,
		Role: existingRole.Name,
	}
	return LoginResponse{Token: token, User: userResponse}, nil
}

func (as *AuthServiceImpl) RegisterUser(req user.UserRequest) error {
	parsedUuid, err := uuid.Parse(req.Role)
	if err != nil {
		return errors.New("invalid uuid")
	}
	_, err = as.rr.GetByUuid(parsedUuid)
	if err != nil {
		return errors.New("role not found")
	}

	isUserExist, err := as.ur.FindUserByEmail(&req.Email)
	if err != nil {
		return err
	}
	if isUserExist != nil {
		return errors.New("user with email " + req.Email + " already exists")
	}

	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	hashedPassword, _ := utils.HashPassword(req.Password)

	user := user.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    hashedPassword,
		RoleUuid:    parsedUuid,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	return as.ar.RegisterUser(&user)
}