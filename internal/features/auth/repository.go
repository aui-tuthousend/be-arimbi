package auth

import (
	"os"
	"time"

	"be-arimbi/internal/features/user"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	// Login(req LoginRequest) (LoginResponse, error)
	CheckPasswordHash(password, hash string) bool
	GenerateJWT(userID uuid.UUID, role string) (string, error)
	RegisterUser(req *user.User) error
	FindUserByEmail(email string) (*user.User, error)
}

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{DB: db}
}

func (ar *AuthRepositoryImpl) CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (ar *AuthRepositoryImpl) GenerateJWT(userID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(time.Hour * 150).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (ar *AuthRepositoryImpl) RegisterUser(req *user.User) error {
	return ar.DB.Create(req).Error
}

func (ar *AuthRepositoryImpl) FindUserByEmail(email string) (*user.User, error) {
	var user user.User
	result := ar.DB.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}