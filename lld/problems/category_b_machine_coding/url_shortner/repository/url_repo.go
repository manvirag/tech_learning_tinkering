package repository

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"url_shortner/models"
)

type IUrlShorterRepo interface {
	SaveShortUrl(longUrl string, shortUrl string) error
	GetLongUrl(urlId string) (error, string)
}

type UrlShorterRepo struct {
	db map[string]models.UrlShorterObj
	sync.RWMutex
}

func NewUrlShorterRepo() IUrlShorterRepo {
	return &UrlShorterRepo{
		db: make(map[string]models.UrlShorterObj),
	}
}

func (usr *UrlShorterRepo) SaveShortUrl(longUrl string, shortUrl string) error {
	usr.Lock()
	defer usr.Unlock()
	if _, ok := usr.db[shortUrl]; ok {
		return errors.New("already exist short url")
	}

	usr.db[shortUrl] = models.UrlShorterObj{
		ShortUrl:     shortUrl,
		LongUrl:      longUrl,
		Expiration:   time.Now().Add(24 * time.Hour),
		CreationTime: time.Now(),
	}
	fmt.Println(usr.db[shortUrl])
	return nil
}
func (usr *UrlShorterRepo) GetLongUrl(urlId string) (error, string) {
	usr.RLock()
	defer usr.RUnlock()
	if urlOb, ok := usr.db[urlId]; !ok {
		return errors.New("not find this url in db"), ""
	} else if urlOb.Expiration.Compare(time.Now()) == -1 {
		fmt.Println(urlOb.Expiration.Second())
		return errors.New("expire "), ""
	} else {
		return nil, urlOb.LongUrl
	}
}
