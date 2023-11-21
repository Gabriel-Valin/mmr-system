package schemas

import "gorm.io/gorm"

type RankedsStatus struct {
	gorm.Model
	PlayerID      int
	Player        Player
	Points        float32
	Mmr           float32
	Md5_completed bool
}
