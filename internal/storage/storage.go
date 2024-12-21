package storage

import "distributed-crawler/internal/models"

type Storage interface {
	Save(url models.URL)
	Exists(url string) bool
}
