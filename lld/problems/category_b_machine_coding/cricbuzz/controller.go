package main

import (
	"fmt"
	"main/models"
	"main/usecase"
)

func main() {
	  teamInd := models.CricketTeam{
			Team : models.Team{
					CountryName: "India",
			},
			Players: []models.Player{ {Name: "Virat Kohli"}, {Name: "Rohit Sharma"}, {Name: "Hardik Pandya"}},
		}
		teamAus := models.CricketTeam{
			Team : models.Team{
					CountryName: "Australia",
			},
			Players: []models.Player{ {Name: "David Warner"}, {Name: "Smit"}, {Name: "C"}},
		}

	  cricketMatch := models.NewCricketMatch("indvsaur",teamInd,teamAus)
	  cricbuzzRepo := models.NewCricbuzzRepo()
	  cricbuzzRepo.AddMatch(cricketMatch)
		
	  // fmt.Printf("%+v", cricbuzzRepo)
	  // match := cricbuzzRepo.GetMatches()
	  // fmt.Printf("%+v", match)
		

		cribuzz_uc := usecase.NewCricbuzzUsecase(cricbuzzRepo)
	  fmt.Printf("%+v", cribuzz_uc.SearchMatches("Live"))
	  
		

	  
}
