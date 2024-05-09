package models

type Chunk struct {
	Identifier string `json:"identifier" gorm:"type:varchar(255);primary_key"`
	Chunk      string `json:"chunk" gorm:"type:varchar(255)"`
	Links      int    `json:"number_of_links" gorm:"type:int"`
}
