package models

type MatchInterface interface {
  GetType() string
  GetMatchId() string
  GetStatus() string
}

type Match struct{
  Type string  
  MatchId string   
}

type CricketMatch struct{
  Match
  TeamA CricketTeam
  TeamB CricketTeam
  ScoreA int
  ScoreB int
  WicketA int
  WicketB int
  Status StatusInterface
}

func NewCricketMatch(matchId string, teamA CricketTeam, teamB CricketTeam) CricketMatch{
   // Todo remove hard coding use constants
    return CricketMatch{
      Match: Match{
        Type: "Cricket",
        MatchId: matchId,
      },
      TeamA:  teamA,
      TeamB:  teamB,
      Status: ScheduledStatus{
        ScheduleTimestamp: 12345,
      },
    }
}

func(cm CricketMatch)GetType() string{
  return cm.Match.Type
}
func(cm CricketMatch)GetMatchId() string{
  return cm.Match.MatchId
}
func(cm CricketMatch)GetStatus() string{
  return cm.Status.GetStatus();
}


