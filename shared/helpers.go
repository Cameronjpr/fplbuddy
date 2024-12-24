package shared

func GetCodeFromClubId(id int) string {
	switch id {
	case 1:
		return "ars" // Arsenal
	case 2:
		return "avl" // Aston Villa
	case 3:
		return "bou" // Bournemouth
	case 4:
		return "bre" // Brentford
	case 5:
		return "bha" // Brighton & Hove Albion
	case 6:
		return "che" // Chelsea
	case 7:
		return "cry" // Crystal Palace
	case 8:
		return "eve" // Everton
	case 9:
		return "ful" // Fulham
	case 10:
		return "ips" // Ipswich Town
	case 11:
		return "lei" // Leicester City
	case 12:
		return "liv" // Liverpool
	case 13:
		return "mci" // Manchester City
	case 14:
		return "mun" // Manchester United
	case 15:
		return "new" // Newcastle United
	case 16:
		return "nfo" // Nottingham Forest
	case 17:
		return "sou" // Southampton
	case 18:
		return "tot" // Tottenham Hotspur
	case 19:
		return "whu" // West Ham United
	case 20:
		return "wol" // Wolverhampton Wanderers
	default:
		return "unknown"
	}
}

var ClubNames = map[int]string{
	1:  "Arsenal",
	2:  "Aston Villa",
	3:  "Bournemouth",
	4:  "Brentford",
	5:  "Brighton & Hove Albion",
	6:  "Chelsea",
	7:  "Crystal Palace",
	8:  "Everton",
	9:  "Fulham",
	10: "Ipswich Town",
	11: "Leicester City",
	12: "Liverpool",
	13: "Manchester City",
	14: "Manchester United",
	15: "Newcastle United",
	16: "Nottingham Forest",
	17: "Southampton",
	18: "Tottenham Hotspur",
	19: "West Ham United",
	20: "Wolverhampton Wanderers",
}
