package servies

import (
	"errors"
	"fmt"
	"url_shortner/repository"
	"url_shortner/utils"
)

type IUrlShorterService interface {
	GetLongUrl(shortId string) (error, string)
	GenerateShortUrl(shortId string) (error, string)
}

type UrlShorterService struct {
	usr     repository.IUrlShorterRepo
	ac      *utils.AtomicCounter
	BaseUrl string
}

func NewUrlShorterService(urlRepo repository.IUrlShorterRepo, ac *utils.AtomicCounter) *UrlShorterService {
	return &UrlShorterService{
		usr:     urlRepo,
		ac:      ac,
		BaseUrl: "http://manvirag.com/api/v1/urlshortner/",
	}
}

func (us *UrlShorterService) GenerateShortUrl(longUrl string) (error, string) {
	// validate request from processing.
	for retry := 0; retry < 3; retry++ {
		shortId := utils.GetBase62(us.ac.NextId())
		fmt.Println(shortId)
		err := us.usr.SaveShortUrl(longUrl, shortId)

		if err != nil {
			fmt.Printf(err.Error()+"retrying.. %d", retry+1)
		} else {
			return nil, us.BaseUrl + shortId
		}
	}

	return errors.New("all retry attempt down, check after some time"), ""

}

func (us *UrlShorterService) GetLongUrl(shortId string) (error, string) {
	return us.usr.GetLongUrl(shortId)
}
