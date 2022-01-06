package model

import (
	"math"
	"sort"
	"time"

	"github.com/manomekun/ddd-sample00/domain/valueobj"
)

type User struct {
	Name       string
	BirthDate  time.Time
	ID         string
	RankScores []valueobj.RankScore
}

func (u *User) RankScoreTotal(targetSeason uint) uint {
	var total uint

	// sort by season asc
	sort.SliceStable(u.RankScores, func(i, j int) bool {
		return u.RankScores[i].Season < u.RankScores[j].Season
	})

	for _, r := range u.RankScores {
		if r.Season > targetSeason {
			continue
		}
		if r.Season == targetSeason {
			total += r.Score
		}
		total += uint(math.Floor(float64(r.Score) * 0.7))
	}
	return total
}

func (u *User) IsValidID() bool {
	return len(u.ID) == 8
}
