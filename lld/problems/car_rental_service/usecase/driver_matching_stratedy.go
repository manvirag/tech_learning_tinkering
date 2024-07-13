package usecase

type DriverMatchingStrategy interface {
	GetMatchingDrivers(source, destination string, uberSystem *Uber) []string
}

type DummyDriverMatchingStrategy struct{}

func (dms DummyDriverMatchingStrategy) GetMatchingDrivers(source, destination string, uberSystem *Uber) []string{
  return []string{"d1234"}
}