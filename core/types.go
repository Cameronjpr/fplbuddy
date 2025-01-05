package core

import "time"

type Group struct {
	Name  string
	Teams []Team
}

type Player struct {
	CreatedAt                *time.Time `json:"created_at"`
	UpdatedAt                *time.Time `json:"updated_at"`
	Id                       int        `json:"id"`
	FplElementId             int        `json:"fpl_element_id"`
	Name                     string     `json:"name"`
	Nickname                 string     `json:"nickname"`
	Form                     float64    `json:"form"`
	ExpectedGoals            float64    `json:"expected_goals"`
	ExpectedAssists          float64    `json:"expected_assists"`
	ExpectedGoalInvolvements float64    `json:"expected_goal_involvements"`
	ExpectedGoalsConceded    float64    `json:"expected_goals_conceded"`
	NowCost                  float64    `json:"now_cost"`
	Position                 int        `json:"position"`
	GwScore                  int        `json:"gw_score"`
	ClubId                   int        `json:"club_id"`
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

type FDREntry struct {
	Event      int `json:"event"`
	TeamId     int `json:"team_id"`
	OpponentId int `json:"opponent_id"`
	Difficulty int `json:"difficulty"`
}

type FDRSchedule map[int][]FDREntry

type Fixture struct {
	Id              int
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	Event           int        `json:"event"`
	KickoffTime     time.Time  `json:"kickoff_time"`
	TeamH           int        `json:"team_h"`
	TeamA           int        `json:"team_a"`
	TeamHDifficulty int        `json:"team_h_difficulty"`
	TeamADifficulty int        `json:"team_a_difficulty"`
}
