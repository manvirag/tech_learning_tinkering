package main

import (
	"fmt"
	"url_shortner/repository"
	"url_shortner/servies"
	"url_shortner/utils"
)

/*
longurl -> getshort -> short url
shorturl -> redirect -> redirect to long url

/api/v1/urlshortner/short?longurl = "" -> short url
/api/v1/urlshortner/?id="" -> redirect to long url

db:

longurl shorturl expire created

expire

how to generate unique id globally -> we already have some solution for this. sonyflake or snowfalke

timestampe serverid, sequenceno., empty

server1

server2         db

server3

coupled, in worst case -> db can down if capacity << load

longurl shorturl expire created status

or async
*/
func main() {
	//fmt.Println(len(utils.BASE62))
	url := "https://www.abc.com/domaina/domainb/pqr"
	//shortUrl :=
	//
	ac := utils.NewAtomicCounter()
	urlRepo := repository.NewUrlShorterRepo()
	urlServ := servies.NewUrlShorterService(urlRepo, ac)
	fmt.Println(len(urlServ.BaseUrl))

	surl := []string{}
	for ix := 0; ix < 10; ix++ {
		//fmt.Println(ac.NextId())
		err, shortUrl := urlServ.GenerateShortUrl(url)
		surl = append(surl, shortUrl)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(shortUrl)
	}
	fmt.Println()
	surl = append(surl, "http://manvirag.com/api/v1/urlshortner/abcdefg")
	surl = append(surl, "http://manvirag.com/api/v1/urlshortner/fdsfsd")
	surl = append(surl, "http://manvirag.com/api/v1/urlshortner/fdsafsdafsadfsdf")

	for ix := 0; ix < 13; ix++ {
		fmt.Println(surl[ix][39:])
		err, longurl := urlServ.GetLongUrl(surl[ix][39:])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(longurl)
	}

}
