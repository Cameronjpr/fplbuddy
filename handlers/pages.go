package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/cameronjpr/fplbuddy/components"
	"github.com/cameronjpr/fplbuddy/core"
	"github.com/cameronjpr/fplbuddy/lib"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	component := components.NotFoundPage()
	component.Render(r.Context(), w)
}

func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	p, err := lib.GetTrendingPlayers()

	if err != nil {
		fmt.Println("error fetching trending players", err)
	}

	response := core.HomepageResponse{
		TrendingPlayers: p,
	}

	component := components.HomePage(response)
	component.Render(r.Context(), w)
}

func HandleTeamPage(w http.ResponseWriter, r *http.Request) {
	teamId, err := r.Cookie("teamId")

	if err != nil {
		fmt.Println("error getting teamId cookie", err)
	}

	fmt.Println(teamId)

	if teamId != nil {
		fmt.Println(teamId.Value)

		http.Redirect(w, r, fmt.Sprintf("/team/%s", teamId.Value), http.StatusMovedPermanently)
	} else {
		component := components.EntryPage(core.Entry{})
		component.Render(r.Context(), w)
	}
}

func HandleComparisonPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	players := []core.Player{}
	wg := sync.WaitGroup{}

	for key, value := range query["id"] {
		fmt.Println(key, value)

		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			playerId, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("error converting player id to int", err)
			}
			p, err := lib.GetPlayer(playerId)
			if err != nil {
				fmt.Println("error fetching player", err)
			}

			players = append(players, p)
		}(value)

	}
	wg.Wait()

	component := components.ComparisonPage(players)
	component.Render(r.Context(), w)

}

func HandlePlayersPage(w http.ResponseWriter, r *http.Request) {
	rawSkip := r.URL.Query().Get("skip")
	rawLimit := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	sort := r.URL.Query().Get("sort")
	dir := r.URL.Query().Get("dir")

	skip, err := strconv.Atoi(rawSkip)
	if err != nil {
		fmt.Println("error converting skip to int", err)
		skip = 0
	}

	limit, err := strconv.Atoi(rawLimit)
	if err != nil {
		fmt.Println("error converting limit to int", err)
		limit = 25
	}

	res, err := lib.GetPlayers(skip, limit, search, sort, dir)

	if err != nil {
		fmt.Println("error fetching players", err)
		http.Error(w, "error fetching players", http.StatusInternalServerError)
	}

	sortOptions := map[string]string{
		"name":     "/players?sort=name",
		"form":     "/players?sort=form",
		"position": "/players?sort=position",
		"now_cost": "/players?sort=now_cost",
	}

	newDir := "asc"
	if dir == "asc" || dir == "" {
		newDir = "desc"
	}
	sortOptions[sort] += fmt.Sprintf("&dir=%s", newDir)

	if r.Header.Get("HX-Request") == "true" {
		component := components.PlayersTable(res, sortOptions)
		component.Render(r.Context(), w)
	} else {
		component := components.PlayersPage(res, sortOptions)
		component.Render(r.Context(), w)
	}
}

func HandleGetEntryPage(w http.ResponseWriter, r *http.Request) {
	teamId := r.PathValue("teamId")
	gw := r.URL.Query().Get("gw")

	if teamId == "" {
		http.Error(w, "Missing Team ID", http.StatusBadRequest)
	}

	gwInt, err := strconv.Atoi(gw)
	if err != nil {
		fmt.Println("error converting gameweek to int", err)
		gwInt = -1
	}

	teamIdInt, err := strconv.Atoi(teamId)
	if err != nil {
		fmt.Println("error converting teamId to int", err)
	}

	entry, err := core.GetEntry(teamIdInt, gwInt)

	if err != nil {
		http.Error(w, "Error fetching entry details", http.StatusInternalServerError)
	}

	component := components.EntryPage(entry)
	component.Render(r.Context(), w)
}

func HandleFDRPage(w http.ResponseWriter, r *http.Request) {
	schedule, err := lib.GetFDRSchedule()

	if err != nil {
		fmt.Println("error fetching fdr schedule", err)
		http.Error(w, "error fetching fdr schedule", http.StatusInternalServerError)
	}

	component := components.FDRPage(schedule)
	component.Render(r.Context(), w)
}

func HandleFixturesPage(w http.ResponseWriter, r *http.Request) {
	fixtures, err := lib.GetFixtures()

	if err != nil {
		fmt.Println("error fetching fixtures", err)
		http.Error(w, "error fetching fixtures", http.StatusInternalServerError)
	}

	component := components.FixturesPage(fixtures)
	component.Render(r.Context(), w)
}
