package models

import (
	"time"

)

// MultiDatasetSearchResponse represents the full API response
type MultiDatasetSearchResponse struct {
	Results []SatelliteImage `json:"results"`
	Meta    ResponseMeta     `json:"meta"`
}

// SatelliteImage represents each image result from the API
type SatelliteImage struct {
	ID                    uint      `gorm:"primaryKey" json:"id"`
	ViewID                string    `gorm:"unique;not null" json:"view_id"`
	SceneID               string    `gorm:"unique;not null" json:"sceneID"`
	Date                  time.Time `gorm:"not null" json:"date"`
	CloudCoverage         float64   `gorm:"not null" json:"cloudCoverage"`
	AwsPath               string    `gorm:"not null" json:"awsPath"`
	Satellite             string    `gorm:"not null" json:"satellite"`
	DataCoveragePercentage float64   `json:"dataCoveragePercentage,omitempty"`
}

// ResponseMeta contains metadata about the API response
type ResponseMeta struct {
	Page  int    `json:"page"`
	Found int    `json:"found"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
}

// TableName sets the table name for GORM
func (SatelliteImage) TableName() string {
	return "satellite_images"
}
