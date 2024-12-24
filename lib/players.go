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

func GetPlayer(elementId int) (core.Player, error) {
	db := shared.LoadEnvVar("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1) // bit weird to do this here
	}
	defer dbpool.Close()

	query := "SELECT * FROM players WHERE fpl_element_id = @elementId"
	rows, err := dbpool.Query(context.Background(), query, pgx.NamedArgs{"elementId": elementId})

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return core.Player{}, err
	}

	player, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[core.Player])

	if err != nil {
		return player, err
	}

	return player, nil
}

// Turn into variadic function?
func GetPlayers(skip, limit int, search string, sort string) (core.PlayersResponse, error) {
	fmt.Printf("searching for players with name: %s\n", search)
	db := shared.LoadEnvVar("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1) // bit weird to do this here
	}
	defer dbpool.Close()

	res := core.PlayersResponse{}

	query := "SELECT * FROM players WHERE name ILIKE @search"
	if sort != "" {
		query += fmt.Sprintf(" ORDER BY %s", sort)
	}
	query += " LIMIT @limit OFFSET @skip"
	rows, err := dbpool.Query(context.Background(), query, pgx.NamedArgs{"search": fmt.Sprintf("%%%s%%", search), "limit": limit, "skip": skip})
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return res, err
	}

	players, err := pgx.CollectRows(rows, pgx.RowToStructByName[core.Player])

	for _, player := range players {
		fmt.Println(player)
	}
	if err != nil {
		return res, err
	}

	res.Players = players

	return res, nil
}

func GetTrendingPlayers() ([]core.Player, error) {
	query := "SELECT * FROM players ORDER BY now_cost DESC LIMIT 12"
	ctx := context.Background()
	db := shared.LoadEnvVar("DATABASE_URL")

	dbpool, err := pgxpool.New(context.Background(), db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1) // bit weird to do this here
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(ctx, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	players, err := pgx.CollectRows(rows, pgx.RowToStructByName[core.Player])
	if err != nil {
		return nil, err
	}
	return players, nil
}
