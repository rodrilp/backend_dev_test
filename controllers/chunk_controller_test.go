package controllers

import (
	"fmt"
	"main/mocks"
	"main/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestChunkCreateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock for repository
	mockChunkRepo := mocks.NewMockChunkRepositoryInterface(ctrl)

	mockChunkRequest := ChunkRequest{
		Identifier: "chunk-11",
		Chunk: []string{"A","T", "G", "C", "C", "A"},
	}

	mockedChunk := models.Chunk{
		Identifier: "chunk-11",
		Chunk: `A,T,G,C,C,A`,
		Links: 2,
	}

	expectedChunk := ChunkResponse{
		Identifier: "chunk-11",
		Chunk: []string{"A","T", "G", "C", "C", "A"},
		Links: 2,
	}

	// Mock Call
	mockChunkRepo.EXPECT().CreateChunk(&mockedChunk).Return(nil)

	processedChunk, err := ProcessCreateChunk(mockChunkRequest, mockChunkRepo)
	
	assert.Nil(t, err)
	assert.Equal(t, expectedChunk, processedChunk)
}

func TestChunkCreateFailiure(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Create mock for repository
	mockChunkRepo := mocks.NewMockChunkRepositoryInterface(ctrl)

	mockChunkRequest := ChunkRequest{
		Identifier: "chunk-11",
		Chunk: []string{"A","T", "G", "C", "C", "A"},
	}

	mockedChunk := models.Chunk{
		Identifier: "chunk-11",
		Chunk: `A,T,G,C,C,A`,
		Links: 2,
	}

	// Mock Call
	mockChunkRepo.EXPECT().CreateChunk(&mockedChunk).Return(fmt.Errorf("Test error handling"))

	_, err := ProcessCreateChunk(mockChunkRequest, mockChunkRepo)
	
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("Test error handling"), err)
}