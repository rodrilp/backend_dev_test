package controllers

import (
	"fmt"
	"main/models"
	"main/repository"
	"main/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type ChunkRequest struct {
	Identifier string   `json:"identifier" binding:"required"`
	Chunk      []string `json:"chunk" binding:"required"`
}

type ChunkResponse struct {
	Identifier string   `json:"identifier"`
	Chunk      []string `json:"chunk"`
	Links      int      `json:"number_of_links"`
}

func ProcessCreateChunk(chunkRequest ChunkRequest, chunkRepository repository.ChunkRepositoryInterface) (ChunkResponse, error) {
	// Call the method to get the number of links inside the chunk
	totalLinks := utils.GetLinks(chunkRequest.Chunk)

	// Convert the arry into string to store in the DB
	stringChunk := strings.Join(chunkRequest.Chunk, ",")

	chunk := models.Chunk{
		Identifier: chunkRequest.Identifier,
		Chunk:      stringChunk,
		Links:      totalLinks,
	}

	// Create Chunk in DB
	err := chunkRepository.CreateChunk(&chunk)

	if err != nil {
		fmt.Printf("Failed to save the chunk in the DB: %s", err.Error())
		return ChunkResponse{}, err
	}

	// Crete the response struct to return
	chunkResponse := ChunkResponse {
		Identifier: chunk.Identifier,
		Chunk: chunkRequest.Chunk,
		Links: chunk.Links,
	}

	return chunkResponse, nil
}

func CreateChunk(ctx *gin.Context) {
	var chunkRequest ChunkRequest

	err := ctx.ShouldBindJSON(&chunkRequest)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid parameters inside the request body"})
		ctx.Abort()
		return
	}
	// Init chunk repository
	chunkRepository := repository.ChunkRepository{}

	// Process the logic
	chunk, err := ProcessCreateChunk(chunkRequest, chunkRepository)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		ctx.Abort()
		return
	}
	
	ctx.JSON(200, chunk)
}
