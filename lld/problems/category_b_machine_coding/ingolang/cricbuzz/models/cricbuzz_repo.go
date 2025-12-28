package models

import ("errors")

type CricbuzzRepoInterface interface{
  GetMatches() []MatchInterface
  GetMatch(matchId string) (MatchInterface, error)
  AddMatch(match MatchInterface) error
  
}

type CricbuzzRepo struct{
  Matches []MatchInterface
}

func NewCricbuzzRepo() *CricbuzzRepo{
  return &CricbuzzRepo{
    Matches: make([]MatchInterface, 0),
  }
}

func (cr *CricbuzzRepo)AddMatch(match MatchInterface) error{
   // Todo: Add check if already exist, match id.
   cr.Matches = append(cr.Matches, match)
   return nil
}

func (cr *CricbuzzRepo)GetMatch(matchId string) (MatchInterface, error){
   for _,match := range cr.Matches{
      if(match.GetMatchId() == matchId){
         return match, nil
      }
   }
   return CricketMatch{}, errors.New("Match not found")
}

func (cr *CricbuzzRepo)GetMatches() []MatchInterface{
   return cr.Matches
}


