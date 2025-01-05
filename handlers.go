package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/cameronjpr/fplbuddy/components"
	"github.com/cameronjpr/fplbuddy/core"
	"github.com/cameronjpr/fplbuddy/lib"
)

func handleSearchPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling search players")

	err := r.ParseForm()

	if err != nil {
		fmt.Println("error parsing form", err)
		http.Error(w, "error parsing form", http.StatusInternalServerError)
		return
	}

	players, err := lib.GetPlayers(-1, -1, "", "", "")

	if err != nil {
		fmt.Println("error fetching players", err)
		http.Error(w, "error fetching players", http.StatusInternalServerError)
		return
	}

	sortOptions := map[string]string{
		"name":     "/players?sort=name&dir=asc",
		"form":     "/players?sort=form&dir=asc",
		"position": "/players?sort=position&dir=asc",
		"now_cost": "/players?sort=now_cost&dir=asc",
	}

	component := components.PlayersTable(players, sortOptions)
	component.Render(r.Context(), w)
}

func handlePostTeam(w http.ResponseWriter, r *http.Request) {
	teamId := r.FormValue("teamId")

	fmt.Printf("teamId: %v\n", teamId)

	for key, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			fmt.Println(key, value)
		}
	}

	if teamId == "" {
		http.Error(w, "Missing Team ID", http.StatusBadRequest)
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:8000/team/%v", teamId))

	if err != nil {
		fmt.Println("error fetching players", err)
		http.Error(w, "error fetching players", http.StatusInternalServerError)
		return
	}

	data := []core.EntryResponse{}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Print("resp", resp.Status)

	if err != nil {
		fmt.Println("error reading body", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("body", body)
	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("error unmarshalling json", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("handling post team", teamId)

	http.Header.Add(w.Header(), "Set-Cookie", fmt.Sprintf("teamId=%s", teamId))
	http.Header.Add(w.Header(), "Location", fmt.Sprintf("/team/%s", teamId))
	http.Header.Add(w.Header(), "HX-Redirect", fmt.Sprintf("/team/%s", teamId))
}

func handlePostComparison(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling post comparison")
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error parsing form, %v\n", err)
	}

	fmt.Println("r.PostForm", r.PostForm)

	query := r.URL.Query()

	for key, values := range r.PostForm {
		if key == "player" {
			for _, value := range values {
				query.Add("id", value)
			}
		}
	}

	if len(query["id"]) < 2 {
		fmt.Println("At least two players are required for comparison")
		http.Error(w, "At least two players are required for comparison", http.StatusBadRequest)
	}

	fmt.Println("handling post comparison with query: ", query)

	http.Header.Add(w.Header(), "Location", fmt.Sprintf("/comparison?%s", query.Encode()))
	http.Header.Add(w.Header(), "HX-Redirect", fmt.Sprintf("/comparison?%s", query.Encode()))
}

func handleGetPlayer(w http.ResponseWriter, r *http.Request) {
	playerId := r.PathValue("playerId")

	if playerId == "" {
		http.Error(w, "Missing Player ID", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(playerId)

	if err != nil {
		fmt.Println("error converting playerId to int", err)
		http.Error(w, "error converting playerId to int", http.StatusInternalServerError)
	}

	p, err := lib.GetPlayer(id)

	if err != nil {
		fmt.Println("error fetching players", err)
		http.Error(w, "error fetching players", http.StatusInternalServerError)
		return
	}

	fdr, err := lib.GetFDRSchedule(p.ClubId)
	if err != nil {
		fmt.Println("error fetching fdr", err)
		http.Error(w, "error fetching fdr", http.StatusInternalServerError)
		return
	}

	http.Header.Add(w.Header(), "Cache-Control", "public, max-age=3600")
	component := components.PlayerPage(p, fdr)
	component.Render(r.Context(), w)
}
