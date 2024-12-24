
package cassandra

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

var cassandraSession *gocql.Session

func InitCassandra(cassandraIP, keyspace string) error {
	cluster := gocql.NewCluster(cassandraIP)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	var err error
	cassandraSession, err = cluster.CreateSession()
	if (err != nil) {
		return err
	}
	return nil
}

func CloseCassandra() {
	if cassandraSession != nil {
		cassandraSession.Close()
	}
}

func UpdateProcessInfo(processID, ip, status string) error {
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

func DeleteProcessInfo(processID string) error {
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