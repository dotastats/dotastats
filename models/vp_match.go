package models

import "time"

type VpMatch struct {
	Id             int
	SeriesID       int
	MatchID        int
	TeamAID        string
	TeamBID        string
	TeamA          string
	TeamB          string
	LogoA          string
	LogoB          string
	TeamAShort     string
	TeamBShort     string
	Tournament     string
	TournamentLogo string
	Game           string
	BestOf         string
	// sub match specific
	URL            string
	Time           *time.Time
	MatchName      string
	ModeName       string
	ModeDesc       string
	HandicapAmount string
	HandicapTeam   string
	RatioA         float64
	RatioB         float64
	Winner         string
	Status         string
	ScoreA         float64
	ScoreB         float64
	Note           string
}
