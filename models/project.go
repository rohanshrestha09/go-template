package models

import (
	"database/sql"

	"github.com/rohanshrestha09/go-template/enums"
)

type Project struct {
	Model
	Name              string              `json:"name" gorm:"not null"`
	Description       string              `json:"description" gorm:"type:text;not null"`
	Published         bool                `json:"published" gorm:"default:true;not null"`
	Image             string              `json:"image"`
	ImageName         string              `json:"imageName"`
	EstimatedDuration uint                `json:"estimatedDuration" gorm:"not null"`
	StartDate         sql.NullTime        `json:"startDate"`
	EndDate           sql.NullTime        `json:"endDate"`
	Status            enums.ProjectStatus `json:"status" gorm:"type:enum('OPEN','IN_PROGRESS','COMPLETED');default:OPEN;not null"`
	UserID            uint                `json:"userId" gorm:"not null"`
	User              *User               `json:"user"`
}
