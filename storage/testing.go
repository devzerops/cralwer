package storage

import (
    "distributed-crawler/models"
    "testing"
)

func TestRedisStorage(t *testing.T) {
    storage := NewRedisStorage()

    url := models.URL{Address: "http://example.com", Priority: 1}
    storage.Save(url)

    if !storage.Exists(url.Address) {
        t.Errorf("Expected URL to exist in Redis")
    }
}