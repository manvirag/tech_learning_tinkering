package util

type AtomicCounter struct {
	count int
}

var ac *AtomicCounter = nil

func NewAtomicCounter() *AtomicCounter {
	if ac != nil {
		return ac
	}
	ac := &AtomicCounter{}
	return ac
}
func (ac *AtomicCounter) GetId() int {
	ac.count = ac.count + 1
	return ac.count
}
