package repository

import (
	"log"
	"os"
	"server-a/config"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	gocqlastra "github.com/datastax/gocql-astra"
)

type Repository struct {
	session *gocql.Session
}

func NewRepository(config *config.Config) *Repository {
	cluster, err := gocqlastra.NewClusterFromBundle(os.Getenv("ASTRA_DB_SECURE_BUNDLE_PATH"),
		"token", os.Getenv("ASTRA_DB_APPLICATION_TOKEN"), 30*time.Second)

	if err != nil {
		log.Panicf("fail to connect cassandra cluster by bundle: %v", err)
	}
	cluster.Timeout = 30 * time.Second
	cluster.Keyspace = config.Cassandra.Keyspace

	session, err := gocql.NewSession(*cluster)

	if err != nil {
		log.Panicf("fail to create session from cassandra cluster: %v", err)
	}

	r := &Repository{session}

	return r
}
