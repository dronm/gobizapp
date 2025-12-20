package models

type TotCount struct {
	TotCount int `json:"tot_count" agg:"count(*)"`
}

