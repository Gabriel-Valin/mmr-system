package services

import (
	"errors"

	"github.com/gabriel-valin/mmr/pkg/schemas"
	"gorm.io/gorm"
)

type CreatePlayerService struct {
	db *gorm.DB
}

func NewCreatePlayerService(db *gorm.DB) *CreatePlayerService {
	return &CreatePlayerService{db: db}
}

func (s *CreatePlayerService) Create(username, password, email string) error {
	player := schemas.Player{}
	err := s.db.First(&player, "username = ?", username).Error
	if err == nil {
		return errors.New("user already registered")
	}
	schema := schemas.Player{
		Username: username,
		Password: password,
		Email:    email,
	}
	return s.db.Create(&schema).Error
}
