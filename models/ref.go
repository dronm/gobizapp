package models

type RefKey struct {
	ID int `json:"id"`
}

type Ref struct {
	Keys     RefKey  `json:"keys"`
	Descr    *string `json:"descr"`
	DataType *string `json:"dataType"`
}

func (r Ref) String() string {
	if r.Descr == nil {
		return ""
	}
	return *r.Descr
}
