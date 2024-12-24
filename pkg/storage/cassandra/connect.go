package cassandra

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

var cassandraSession *gocql.Session
var cassandraIP string

func init() {
	if _, err := os.Stat(".env"); err == nil {
		if os.Getenv("EXTERNAL_DB") == "true" {
			cassandraIP = os.Getenv("DB_HOST_ADDR")
		} else {
			cassandraIP = "127.0.0.1"
		}
	} else {
		cassandraIP = "127.0.0.1"
	}
}

type CassandraStorage struct{}

func (cs *CassandraStorage) Init(keyspace string, overrideIP ...string) error {
	ip := cassandraIP
	if len(overrideIP) > 0 {
		ip = overrideIP[0]
	}

	cluster := gocql.NewCluster(ip)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	var err error
	cassandraSession, err = cluster.CreateSession()
	if err != nil {
		return err
	}
	return nil
}

func (cs *CassandraStorage) Close() {
	if cassandraSession != nil {
		cassandraSession.Close()
	}
}

func (cs *CassandraStorage) UpdateProcessInfo(processID, ip, status string) error {
	if cassandraSession == nil {
		return fmt.Errorf("cassandra session is not initialized")
	}

	query := `INSERT INTO processes (process_id, ip, status, updated_at) VALUES (?, ?, ?, ?)`
	err := cassandraSession.Query(query, processID, ip, status, time.Now()).Exec()
	if err != nil {
		log.Printf("Failed to update process info: %v", err)
		return err
	}
	return nil
}

func (cs *CassandraStorage) DeleteProcessInfo(processID string) error {
	if cassandraSession == nil {
		return fmt.Errorf("cassandra session is not initialized")
	}

	query := `DELETE FROM processes WHERE process_id = ?`
	err := cassandraSession.Query(query, processID).Exec()
	if err != nil {
		log.Printf("Failed to delete process info: %v", err)
		return err
	}
	return nil
}