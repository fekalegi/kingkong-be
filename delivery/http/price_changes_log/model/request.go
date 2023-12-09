package model

import "time"

type List struct {
	Limit     int        `form:"limit" validate:"required,min=-1"`
	Offset    int        `form:"offset" validate:"min=0"`
	Search    string     `form:"search"`
	StartDate *time.Time `form:"start_date"`
	EndDate   *time.Time `form:"end_date"`
}
