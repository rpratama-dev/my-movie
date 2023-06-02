package models

import "github.com/google/uuid"

type Answer struct {
	BaseModelID
	Content				string
	IsTheBest			bool
	QuestionID  	uuid.UUID	`gorm:"type:uuid;index;not null" json:"question_id"`
	Question    	Question	`gorm:"foreignKey:QuestionID" json:"question"`
	UserID      	uuid.UUID	`gorm:"type:uuid;index;not null" json:"user_id"`
	User        	User      `gorm:"foreignKey:UserID" json:"user"`
	BaseModelAudit
}