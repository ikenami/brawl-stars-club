package bservice

import (
	"net/http"
	"fmt"
	"sort"

	bgateway "brawlstarsclub.com/m/v2/gateway/brawlstars"
	"brawlstarsclub.com/m/v2/entity"
)




func GetClub(w http.ResponseWriter, req *http.Request) {
	club, err := bgateway.GetClub()
	if err != nil {
		panic (err)
	}

	result := []entity.RaceMember{}

	for _, member := range club.Members {
		player, err := bgateway.GetPlayer(member.Tag)
		if err != nil {
			panic(err)
		}

		memberInitTrophies := entity.InitTrophies[member.Tag]
		deltaTrophies := player.Trophies - memberInitTrophies

		raceMember := entity.RaceMember{
			Name: player.Name,
			Tag: member.Tag,
			InitTrophies: memberInitTrophies,
			CurrentTrophies: player.Trophies,
			DeltaTrophies: deltaTrophies,
		}
		result = append(result, raceMember)
	}

	sort.Sort(entity.Leaderboard(result))
	printLeaderboard(result)
}

func printLeaderboard(players []entity.RaceMember) {
	for i, player := range players {
		fmt.Printf("%v ) %v : %v troféis (atual: %v)\n", i + 1, player.Name, player.DeltaTrophies, player.CurrentTrophies)
	}
}
