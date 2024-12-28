package core

import "time"

type Group struct {
	Name  string
	Teams []Team
}

type Player struct {
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Id           int        `json:"id"`
	FplElementId int        `json:"fpl_element_id"`
	Name         string     `json:"name"`
	Nickname     string     `json:"nickname"`
	Form         float64    `json:"form"`
	NowCost      float64    `json:"now_cost"`
	Position     int        `json:"position"`
	GwScore      int        `json:"gw_score"`
	ClubId       int        `json:"club_id"`
}

type EntryPlayer struct {
	Id           int     `json:"id"`
	FplElementId int     `json:"fpl_element_id"`
	Name         string  `json:"name"`
	Nickname     string  `json:"nickname"`
	Form         float64 `json:"form"`
	NowCost      float64 `json:"now_cost"`
	Position     int     `json:"position"`
	Location     int     `json:"location"`
	GwScore      int     `json:"gw_score"`
	ClubId       int     `json:"club_id"`
}

type Entry struct {
	Id          string
	Gameweek    int
	Total       int
	Bench       []EntryPlayer
	Goalkeepers []EntryPlayer
	Defenders   []EntryPlayer
	Midfielders []EntryPlayer
	Forwards    []EntryPlayer
}

type Team struct {
	Name string
}

type HomepageResponse struct {
	TrendingPlayers []Player `json:"players"`
}

type EntryResponse struct {
	Gameweek int           `json:"gw"`
	Players  []EntryPlayer `json:"players"`
}

type PlayersResponse struct {
	Players []Player `json:"players"`
	Index   int      `json:"index"`
	Total   int      `json:"total"`
	Next    int      `json:"next"`
	Prev    int      `json:"prev"`
}
