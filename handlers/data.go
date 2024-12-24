package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cameronjpr/fplbuddy/components"
	"github.com/cameronjpr/fplbuddy/core"
	"github.com/cameronjpr/fplbuddy/lib"
)

func HandleGetEntryView(w http.ResponseWriter, r *http.Request) {
	teamId := r.PathValue("teamId")
	gw := r.URL.Query().Get("gw")

	if teamId == "" {
		http.Error(w, "Missing Team ID", http.StatusBadRequest)
	}

	gwInt, err := strconv.Atoi(gw)
	if err != nil {
		fmt.Println("error converting gameweek to int", err)
		gwInt = 1
	}

	teamIdInt, err := strconv.Atoi(teamId)
	if err != nil {
		fmt.Println("error converting teamId to int", err)
	}

	entry, err := core.GetEntry(teamIdInt, gwInt)

	if err != nil {
		http.Error(w, "Error fetching entry details", http.StatusInternalServerError)
	}

	component := components.TeamEntry(entry)
	component.Render(r.Context(), w)
}

func HandleSearchPlayers(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println("error parsing form", err)
		http.Error(w, "error parsing form", http.StatusInternalServerError)
		return
	}

	search := r.FormValue("search")

	if search == "" {
		component := components.PlayerSearchBar([]core.Player{})
		component.Render(r.Context(), w)
	}

	result, err := lib.GetPlayers(0, 5, search, "")

	if err != nil {
		fmt.Println("error fetching players", err)
		http.Error(w, "error fetching players", http.StatusInternalServerError)
		return
	}

	component := components.PlayerSearchBar(result.Players)
	component.Render(r.Context(), w)
}
