package usecase

type FareStrategy interface{
  GetFare(source, destination string) int32
}

type DummyFare struct{}

func (df DummyFare)GetFare(source, destination string) int32{
  return 50
}