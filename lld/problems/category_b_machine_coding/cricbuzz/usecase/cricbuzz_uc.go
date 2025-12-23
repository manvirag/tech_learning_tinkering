package usecase

import "main/models"

type CricbuzzUsecaseInterace interface{
  GetMatches() []models.MatchInterface
  SearchMatches(status string) []models.MatchInterface
  GetMatchInfo(matchId string) ( models.MatchInterface, error )
}

type CricbuzzUsecase struct{
  cricbuzzRepo *models.CricbuzzRepo 
}

func NewCricbuzzUsecase(cricbuzzRepo *models.CricbuzzRepo) * CricbuzzUsecase{
  return &CricbuzzUsecase{cricbuzzRepo:cricbuzzRepo}
}

func (c *CricbuzzUsecase) GetMatches() []models.MatchInterface {
  return c.cricbuzzRepo.GetMatches()
}

func (c *CricbuzzUsecase) SearchMatches(status string) []models.MatchInterface {
  matches := c.cricbuzzRepo.GetMatches()
  filteredMatches := []models.MatchInterface{}
  for _, match := range matches {
    if match.GetStatus() == status {
      filteredMatches = append(filteredMatches, match)
    }
  }
  return filteredMatches
}

func (c *CricbuzzUsecase) GetMatchInfo(matchId string) (models.MatchInterface, error) {
  return c.cricbuzzRepo.GetMatch(matchId)
  
}
