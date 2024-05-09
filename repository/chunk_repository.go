package repository

import (
	"main/config"
	"main/models"
)

type ChunkRepositoryInterface interface {
	CreateChunk(chunk *models.Chunk) error
}

type ChunkRepository struct{}

func (ChunkRepository) CreateChunk(chunk *models.Chunk) error {
	db := config.GetDB()

	err := db.Create(&chunk).Error

	return err
}
