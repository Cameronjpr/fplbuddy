package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetEntry(teamId, gw int) (Entry, error) {
	fmt.Printf("fetching entry %v, %v\n", teamId, gw)
	resp, err := http.Get(fmt.Sprintf("http://localhost:8000/team/%v?gw=%v", teamId, gw))

	if err != nil {
		fmt.Println("error fetching players", err)
		return Entry{}, err
	}

	data := EntryResponse{}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error reading body", err)
		return Entry{}, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("error unmarshalling json", err)
		return Entry{}, err
	}

	strId := strconv.Itoa(teamId)

	entry := Entry{
		Id:          strId,
		Gameweek:    data.Gameweek,
		Total:       0,
		Bench:       []EntryPlayer{},
		Goalkeepers: []EntryPlayer{},
		Defenders:   []EntryPlayer{},
		Midfielders: []EntryPlayer{},
		Forwards:    []EntryPlayer{},
	}

	fmt.Println(data)

	for _, player := range data.Players {
		fmt.Println(player)
		if player.Location > 11 {
			entry.Bench = append(entry.Bench, player)
		} else {
			switch player.Position {
			case 1:
				entry.Goalkeepers = append(entry.Goalkeepers, player)
			case 2:
				entry.Defenders = append(entry.Defenders, player)
			case 3:
				entry.Midfielders = append(entry.Midfielders, player)
			case 4:
				entry.Forwards = append(entry.Forwards, player)
			}

			entry.Total += player.GwScore
		}
	}

	fmt.Println(entry)

	return entry, nil
}
