package schemas

import (
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	A_team_players []Player `gorm:"serializer:json"`
	B_team_players []Player `gorm:"serializer:json"`
	A_team_mmr     float32
	B_team_mmr     float32
	Team_won       string
	Mvp            string
}
