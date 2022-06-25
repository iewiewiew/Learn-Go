package model

import "time"

/**
 * @author       weimenghua
 * @time         2024-07-20 14:07
 * @description
 */

type Example struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Num       int        `json:"num"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (*Example) TableName() string {
	return "example"
}
