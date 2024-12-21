package storage

import "distributed-crawler/models"

type Storage interface {
	Save(url models.URL)
	Exists(url string) bool
}