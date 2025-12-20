package models

import (
	fields "github.com/dronm/crudifier/fields"
)

const (
	clientComponentRelation = "client_components"
)

type ClientComponent struct {
	ID                       fields.FieldInt  `json:"id" srvCalc:"true" primaryKey:"true"`
	Caption                  fields.FieldText `json:"caption" required:"TRUE" alias:"Наименование" max:"250"`
	ClientComponentSectionID fields.FieldInt  `json:"client_component_section_id" required:"TRUE" alias:"Секция"`
	Name                     fields.FieldText `json:"name" required:"TRUE" alias:"Vue name" max:"250"`
	Path                     fields.FieldText `json:"path" required:"TRUE" alias:"Vue path"`
	Component                fields.FieldText `json:"component" required:"TRUE" alias:"Vue component"`
	CommentText              fields.FieldText `json:"comment_text"`
}

func (m ClientComponent) Relation() string {
	return clientComponentRelation
}

func (m ClientComponent) CollectionAgg() any {
	return &TotCount{0}
}

type ClientComponentKey struct {
	ID fields.FieldInt `json:"id" required:"true"`
}

func (m ClientComponentKey) Relation() string {
	return clientComponentRelation
}

type ClientComponentUpdate struct {
	Keys  ClientComponentKey `json:"keys"`
	Model ClientComponent    `json:"model"`
}

type ClientComponentDelete struct {
	Keys []ClientComponentKey `json:"keys"`
}
