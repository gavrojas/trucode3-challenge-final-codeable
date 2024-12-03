package models

import "gorm.io/gorm"

type Filters struct {
	gorm.Model    `json:"-"`
	UserID        uint   `json:"user_id"`
	Education     string `json:"education"`
	MaritalStatus string `json:"marital_status"`
	Occupation    string `json:"occupation"`
	MinAge        int    `json:"min_age"`
	MaxAge        int    `json:"max_age"`
	Income        string `json:"income"`
}
