package repository

import (
	"log"
	"server-a/config"

	"github.com/redis/rueidis"
)

type Repository struct {
	client rueidis.Client
}

func NewRepository(config *config.Config) *Repository {
	r := new(Repository)
	c := rueidis.ClientOption{
		InitAddress: config.Redis.URLS,
	}
	var err error
	r.client, err = rueidis.NewClient(c)
	if err != nil {
		log.Panicf("Fail to connect to redis: %v", err)
	}

	return r
}
