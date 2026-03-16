package models

import (
	"fmt"
)

type FinancialMetric struct {
	Value float64
}

func GetFinancialMetrics() []FinancialMetric {
	return []FinancialMetric{{Value: 10.5}, {Value: 20.8}}
}