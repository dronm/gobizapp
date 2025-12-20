package models

import "github.com/dronm/crudifier/fields"

const (
	constantRelation = "constants_list"
)

type ConstantList struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Descr       string  `json:"descr"`
	Val         *string `json:"val"`
	CtrlClass   *string `json:"ctrl_class"`
	CtrlOptions *string `json:"ctrl_options"`
	ViewClass   *string `json:"view_class"`
	ViewOptions *string `json:"view_options"`
}

func (m ConstantList) Relation() string {
	return constantRelation
}

func (m ConstantList) CollectionAgg() any {
	return &TotCount{0}
}

type ConstantValList struct {
	ID          string  `json:"id"`
	Val         any `json:"val"` //*string
}

type ConstantSet struct {
	Val fields.FieldText `json:"val" required:"true" maxLen:"10000"`
}

