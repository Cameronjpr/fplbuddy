package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/cameronjpr/fplbuddy/core"
	"github.com/cameronjpr/fplbuddy/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetFDRSchedule() (core.FDRSchedule, error) {
	db := shared.LoadEnvVar("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1) // bit weird to do this here
	}
	defer dbpool.Close()

	query := "SELECT * FROM fixtures"
	rows, err := dbpool.Query(context.Background(), query)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return core.FDRSchedule{}, err
	}

	fixtures, err := pgx.CollectRows(rows, pgx.RowToStructByName[core.Fixture])

	if err != nil {
		return core.FDRSchedule{}, err
	}

	schedule := make(map[int][]core.FDREntry)

	for _, fixture := range fixtures {
		for i := 0; i < 20; i++ {
			if fixture.TeamH == i {
				schedule[i] = append(schedule[i], core.FDREntry{
					Event:      fixture.Event,
					TeamId:     fixture.TeamH,
					OpponentId: fixture.TeamA,
					Difficulty: fixture.TeamHDifficulty,
				})
			} else if fixture.TeamA == i {
				schedule[i] = append(schedule[i], core.FDREntry{
					Event:      fixture.Event,
					TeamId:     fixture.TeamA,
					OpponentId: fixture.TeamH,
					Difficulty: fixture.TeamADifficulty,
				})
			}
		}
	}

	return schedule, nil
}
