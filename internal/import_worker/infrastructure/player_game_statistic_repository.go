package infrastructure

import (
	"context"
	"database/sql"
	"pointfive/internal/config"
	"pointfive/internal/import_worker/domain"
)

type PlayerGameStatisticRepository struct {
	db *sql.DB
}

func NewPlayerGameStatisticRepository(db *sql.DB, config *config.Config) domain.ImportJobFileRepository {
	r := PlayerGameStatisticRepository{
		db: db,
	}
	return &r
}

func (c *PlayerGameStatisticRepository) AddPlayerGameStatistics(ctx context.Context, playerGameStatistics []domain.PlayerGameStatistic, jobID int) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(`
	INSERT INTO player_stats_raw (
		 job_id, season_id, game_id, team_id, player_id,
		points, rebounds, assists, steals, blocks, fouls,
		turnovers, minutes_played
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, stat := range playerGameStatistics {
		_, err := stmt.Exec(
			jobID, stat.SeasonID, stat.GameID, stat.TeamID, stat.PlayerID,
			stat.Points, stat.Rebounds, stat.Assists, stat.Steals, stat.Blocks, stat.Fouls,
			stat.Turnovers, stat.MinutesPlayed,
		)
		if err != nil {
			return err
		}

	}
	_, err = tx.Exec(`UPDATE import_job_files SET status = $1 WHERE id = $2`, "insert", jobID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}

func (c *PlayerGameStatisticRepository) Get(ctx context.Context, id int) (*domain.ImportJobFile, error) {

	var importJobFile domain.ImportJobFile
	err := c.db.QueryRow(`SELECT id, path, time, status FROM import_job_files WHERE id = $1`, id).
		Scan(&importJobFile.JobID, &importJobFile.Path, &importJobFile.Time, &importJobFile.Status)
	if err != nil {

		return nil, err
	}
	return &importJobFile, nil
}
