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

func GetFixtures() ([]core.Fixture, error) {
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
		return []core.Fixture{}, err
	}

	fixtures, err := pgx.CollectRows(rows, pgx.RowToStructByName[core.Fixture])
	if err != nil {
		return []core.Fixture{}, err
	}

	return fixtures, nil
}
