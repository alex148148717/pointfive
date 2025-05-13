package domain

import "time"

type ImportJobFile struct {
	JobID  int
	Path   string
	Time   time.Time
	Status string
}
type WorkerJob struct {
	ID    string
	RunID string
}
type PlayerGameStatistic struct {
	SeasonID      int32   `json:"season_year"`
	GameID        int32   `json:"game_id"`
	TeamID        int32   `json:"team_id"`
	PlayerID      int32   `json:"player_id"`
	Points        int16   `json:"points"`
	Rebounds      int16   `json:"rebounds"`
	Assists       int16   `json:"assists"`
	Steals        int16   `json:"steals"`
	Blocks        int16   `json:"blocks"`
	Fouls         uint8   `json:"fouls"`
	Turnovers     int16   `json:"turnovers"`
	MinutesPlayed float32 `json:"minutes_played"`
}
