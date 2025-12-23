package models

import "time"

type UrlShorterObj struct {
	LongUrl      string
	ShortUrl     string
	Expiration   time.Time
	CreationTime time.Time
}
