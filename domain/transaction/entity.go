package transaction

import (
	"kingkong-be/common"
	"kingkong-be/helper"
	"log"
	"time"
)

type Transaction struct {
	TransactionID         int        `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID                int        `json:"user_id"`
	CustomerID            int        `json:"customer_id"`
	TransactionType       string     `json:"transaction_type"`
	TransactionDate       *time.Time `json:"transaction_date"`
	TotalPrice            float64    `json:"total_price"`
	AdditionalInformation string     `json:"additional_information"`
	CreatedDate           *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate           *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

type TransactionPart struct {
	TransactionPartID int        `json:"transaction_part_id" gorm:"primaryKey;autoIncrement"`
	TransactionID     int        `json:"transaction_id"`
	PartID            int        `json:"part_id"`
	PartName          string     `json:"part_name" gorm:"->"`
	Quantity          int        `json:"quantity"`
	Price             float64    `json:"price"`
	CreatedDate       *time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate       *time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

type TransactionReport struct {
	TransactionID         int               `json:"transaction_id" gorm:"primaryKey;autoIncrement"`
	UserID                int               `json:"user_id"`
	Username              string            `json:"username"`
	CustomerID            int               `json:"customer_id"`
	CustomerName          string            `json:"customer_name"`
	TransactionType       string            `json:"transaction_type"`
	TransactionDate       *time.Time        `json:"transaction_date"`
	TotalPrice            float64           `json:"total_price"`
	AdditionalInformation string            `json:"additional_information"`
	TransactionParts      []TransactionPart `json:"transaction_parts" gorm:"foreignKey:transaction_id"`
	CreatedDate           *time.Time        `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate           *time.Time        `json:"updated_date" gorm:"autoUpdateTime"`
}

type RequestInsertTransaction struct {
	Transaction      Transaction       `json:"transaction"`
	TransactionParts []TransactionPart `json:"transaction_parts"`
}

type ResponseChart struct {
	WeeklyChartSales     []WeeklyChart  `json:"weekly_chart_sales"`
	WeeklyChartPurchase  []WeeklyChart  `json:"weekly_chart_purchase"`
	MonthlyChartSales    []MonthlyChart `json:"monthly_chart_sales"`
	MonthlyChartPurchase []MonthlyChart `json:"monthly_chart_purchase"`
}

type WeeklyChart struct {
	DayOfWeek int     `json:"day_of_week"`
	Day       string  `json:"day"`
	Sum       float64 `json:"sum"`
}

type MonthlyChart struct {
	Month    int     `json:"month"`
	MonthStr string  `json:"month_str"`
	Sum      float64 `json:"sum"`
}

func (tr *TransactionReport) TableName() string {
	return "transactions"
}

func generateEmptyWeeklyChart(wc []WeeklyChart) []WeeklyChart {
	var res []WeeklyChart
	m := make(map[int]float64)
	for _, v := range wc {
		m[v.DayOfWeek] = v.Sum
	}
	log.Printf("%+v", m)

	for i := 7; i >= 1; i-- {
		now := time.Now()
		sum := 0.00

		if i > 1 {
			now = now.AddDate(0, 0, (i-1)*-1)
		}

		_, ok := m[helper.GetWeekdays(now)]
		if ok {
			sum = m[helper.GetWeekdays(now)]
		}

		res = append(res, WeeklyChart{
			DayOfWeek: int(now.Weekday()),
			Day:       common.Days[int(now.Weekday())],
			Sum:       sum,
		})
	}

	return res
}

func generateEmptyMonthlyChart(wc []MonthlyChart) []MonthlyChart {
	var res []MonthlyChart
	m := make(map[int]float64)
	for _, v := range wc {
		m[v.Month] = v.Sum
		log.Printf("KOKOK %+v", v)
	}

	now := time.Now().Month()
	for i := 11; i >= 0; i-- {
		month := int(now) - i
		sum := 0.00
		log.Println(month, int(now), i)

		if month <= 0 {
			month += 12
		}
		log.Println("KEDUA", month)

		_, ok := m[month]
		if ok {
			sum = m[month]
		}

		res = append(res, MonthlyChart{
			Month:    month,
			MonthStr: common.Month[int(month)],
			Sum:      sum,
		})
	}

	return res
}
