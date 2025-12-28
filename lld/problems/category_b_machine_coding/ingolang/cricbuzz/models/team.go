package models

type TeamInterface interface {
  GetCountryName() string
}

type Team struct {
  CountryName string
}


type CricketTeam struct {
  Team
  Players []Player  
}

