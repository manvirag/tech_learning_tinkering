package models

import (
	"sync"
	"time"
)

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type User struct {
	UserId  int
	Name    string
	GenderV Gender
	Age     int
	// other details
}
type Advertiser struct {
	AdvertiserId int
	Name         string
	Budget       int
	// other details

}

type Campaign struct {
	Id           int
	Content      string
	MaxAge       int
	MinAge       int
	Gender       []Gender
	AdvertiserId int
	BidAmount    int
	// other details
}

type Advertisement struct {
	AdvertisementId int
	CampaignDetail  Campaign
}

type GlobalAdRateLimitI interface {
	ShouldPass() bool
}

type BucketingRateLimiterPerMin struct {
	MaxRequests     int
	TotalRequests   int
	LastUpdatedTime time.Time
}

func (br *BucketingRateLimiterPerMin) ShouldPass() bool {
	diff := time.Now().Minute() - br.LastUpdatedTime.Minute()
	if diff >= 1 {
		br.TotalRequests = 0
	}
	if br.TotalRequests >= br.MaxRequests {
		return false
	}
	br.TotalRequests = br.TotalRequests + 1
	br.LastUpdatedTime = time.Now()
	return true
}

type AtomicCounter struct {
	Count int
	sync.Mutex
}

func (ac *AtomicCounter) GetId() int {
	ac.Lock()
	defer ac.Unlock()
	ac.Count = ac.Count + 1
	return ac.Count
}
