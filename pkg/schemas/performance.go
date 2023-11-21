package schemas

import "gorm.io/gorm"

type Perfomance struct {
	gorm.Model
	PlayerID int
	Player   Player
	MatchID  int
	Match    Match
	Kd       float32
	Kda      float32
	Won      bool
	Points   int
}
