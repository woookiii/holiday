package repository

import (
	"log"
	"server-a/config"
	"time"

	"github.com/apache/cassandra-gocql-driver/v2"
	"github.com/apache/cassandra-gocql-driver/v2/lz4"
)

type Repository struct {
	session *gocql.Session
	rtExp   int64
}

func NewRepository(config *config.Config) *Repository {

	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "default"
	cluster.Timeout = 1 * time.Minute
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = config.Cassandra.Keyspace
	cluster.Compressor = &lz4.LZ4Compressor{}
	//cluster.PageSize = 1000
	//cluster.NextPagePrefetch = 0.25
	//cluster.Tracer =

	session, err := gocql.NewSession(*cluster)

	if err != nil {
		log.Panicf("fail to create session from cassandra cluster: %v", err)
	}

	r := &Repository{
		session: session,
		rtExp:   config.Exp.RtExp,
	}

	return r
}
