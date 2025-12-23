package models

type StatusInterface interface {
	GetStatus() string
}



type ScheduledStatus struct {
	ScheduleTimestamp int64
}

type PastStatus struct {
  
}

type LiveStatus struct {

}

func (ss ScheduledStatus)GetStatus() string{
  return "Scheduled"
}
func (ss PastStatus)GetStatus() string{
  return "Past"
}
func (ss LiveStatus)GetStatus() string{
  return "Live"
}