package models

type Collection struct {
	Data any `json:"data"`
	Agg  any `json:"agg"`
}

type CollectionAlterResult struct {
	AffectedRows int64 `json:"affected_rows"`
}
