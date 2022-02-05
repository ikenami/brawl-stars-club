package entity

type GetClubResponse struct {
	Tag string `json:"tag"`
	Name string `json:"name"`
	Description string `json:"description"`
	ClubType string `json:"type"`
	RequiredTrophies int64 `json:"requiredTrophies"`
	Trophies int64 `json:"trophies"`
	Members []ClubMember `json:"members"`
}

type ClubMember struct {
	Tag string `json:"tag"`
	Name string `json:"name"`
	NameColor string `json:"nameColor"`
	Role string `json:"role"`
	Trophies int64 `json:"trophies"`
}

type GetPlayerResponse struct {
	Tag string `json:"tag"`
	Name string `json:"name"`
	NameColor string `json:"nameColor"`
	Trophies int64 `json:"trophies"`
	HighestTrophies int64 `json:"highestTrophies"`
	ExpLevel int64 `json:"expLevel"`
	TeamVictories int64 `json:"3vs3Victories"`
	SoloVictories int64 `json:"soloVictories"`
	DuoVictories int64 `json:"duoVictories"`
	BestRoboRumbleTime int64 `json:"bestRoboRumbleTime"`
	BestTimeAsBigBrawler int64 `json:"bestTimeAsBigBrawler"`
}

type RaceMember struct {
	Name string
	Tag string
	InitTrophies int64
	CurrentTrophies int64
	DeltaTrophies int64
}

type Leaderboard []RaceMember

func (s Leaderboard) Len() int {
 return len(s)
}

func (s Leaderboard) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Leaderboard) Less(i, j int) bool {
	return s[i].DeltaTrophies > s[j].DeltaTrophies
}