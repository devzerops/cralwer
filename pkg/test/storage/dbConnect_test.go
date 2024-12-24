package test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var hostAddr = "155.248.164.28"

// CassandraConnection tests the connection to Cassandra.
func TestCassandraConnection(t *testing.T) {
	cluster := gocql.NewCluster(hostAddr)
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		t.Fatalf("Failed to connect to Cassandra: %v", err)
	}
	defer session.Close()

	if err := session.Query("SELECT now() FROM system.local").Exec(); err != nil {
		t.Fatalf("Failed to execute query on Cassandra: %v", err)
	}
}

// RedisConnection tests the connection to Redis.
func TestRedisConnection(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: hostAddr + ":6379",
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// MongoDBConnection tests the connection to MongoDB.
func TestMongoDBConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+hostAddr+":27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			t.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()
}
// KafkaConnection tests the connection to Kafka.
func TestKafkaConnection(t *testing.T) {
	conn, err := kafka.Dial("tcp", hostAddr+":9092")
	if err != nil {
		t.Fatalf("Failed to connect to Kafka: %v", err)
	}
	defer conn.Close()
}
