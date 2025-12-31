package service

import (
	. "advertisement/models"
	"errors"
	"fmt"
)

type AdMatchingI interface {
	SuggestAdvertisement(ae *AdvertisementEngine, username string) (Advertisement, error)
}

type AttributeMatching struct {
}

func (am AttributeMatching) SuggestAdvertisement(ae *AdvertisementEngine, username string) (Advertisement, error) {
	// some ml related stuff
	ads := []Advertisement{}
	if usr, ok := ae.UserMap[username]; !ok {
		return Advertisement{}, errors.New("register user first")
	} else {
		for _, ad := range ae.CampaignStore {
			if usr.Age >= ad.MinAge && usr.Age <= ad.MaxAge {
				if _, ok := ae.UserAdvertisementCount[usr.UserId]; !ok {
					ae.UserAdvertisementCount[usr.UserId] = make(map[int]int)
				}
				if _, ok := ae.UserAdvertisementCount[usr.UserId][ad.Id]; !ok {
					ae.UserAdvertisementCount[usr.UserId][ad.Id] = 0
				}
				fetchAddCount := ae.UserAdvertisementCount[usr.UserId][ad.Id]
				if fetchAddCount >= ae.UserFetchCount {
					continue
				}
				ads = append(ads, Advertisement{
					AdvertisementId: ae.ac.GetId(),
					CampaignDetail:  ad,
				})
			}
		}
	}
	if len(ads) > 0 {
		return ads[0], nil
	}
	return Advertisement{}, errors.New("not found")
}

// assuming name as unique for simplicity
type AdvertisementEngine struct {
	AdvertiserMap          map[string]Advertiser
	UserMap                map[string]User
	CampaignStore          map[int]Campaign
	ac                     *AtomicCounter
	AdMatchingStrategy     AdMatchingI
	UserAdvertisementCount map[int]map[int]int
	RateLimiter            GlobalAdRateLimitI
	UserFetchCount         int
}

func NewAdvertisementEngine() *AdvertisementEngine {
	return &AdvertisementEngine{
		AdvertiserMap:          make(map[string]Advertiser),
		UserMap:                make(map[string]User),
		ac:                     &AtomicCounter{},
		CampaignStore:          make(map[int]Campaign),
		AdMatchingStrategy:     AttributeMatching{},
		UserAdvertisementCount: make(map[int]map[int]int),
		RateLimiter:            &BucketingRateLimiterPerMin{MaxRequests: 5},
		UserFetchCount:         10,
	}
}

func (ae *AdvertisementEngine) AddAdvertiser(name string) error {
	if _, ok := ae.AdvertiserMap[name]; ok {
		return errors.New("already exist")
	}
	av := Advertiser{
		AdvertiserId: ae.ac.GetId(),
		Name:         name,
	}
	ae.AdvertiserMap[name] = av
	fmt.Println(ae.AdvertiserMap)
	return nil
}

func (ae *AdvertisementEngine) AddAdvertiserBudget(name string, budget int) error {
	if av, ok := ae.AdvertiserMap[name]; !ok {
		return errors.New("not exist exist")
	} else {
		av.Budget = av.Budget + budget
		ae.AdvertiserMap[name] = av
		fmt.Println(ae.AdvertiserMap)
	}
	return nil
}

func (ae *AdvertisementEngine) AddUser(name string, age int, gender Gender) error {
	if _, ok := ae.UserMap[name]; ok {
		return errors.New("already exist")
	}
	usr := User{
		UserId:  ae.ac.GetId(),
		Name:    name,
		Age:     age,
		GenderV: gender,
	}
	ae.UserMap[name] = usr
	fmt.Println(ae.UserMap)
	return nil
}

func (ae *AdvertisementEngine) CreateCampaign(advertiserId int, bidAmount int, content string, maxAge int, minAge int, genders []Gender) {
	c := Campaign{
		Id:           ae.ac.GetId(),
		Content:      content,
		MaxAge:       maxAge,
		MinAge:       minAge,
		Gender:       genders,
		AdvertiserId: advertiserId,
		BidAmount:    bidAmount,
	}

	ae.CampaignStore[c.Id] = c
}

func (ae *AdvertisementEngine) MatchAdvertisement(username string) (Advertisement, error) {
	if ae.RateLimiter.ShouldPass() == false {
		return Advertisement{}, errors.New("try after some time")
	}
	return ae.AdMatchingStrategy.SuggestAdvertisement(ae, username)
}
