package models

import (
	fields "github.com/dronm/crudifier/fields"
)

const (
	clientComponentSectionRelation = "client_component_sections"
)

type ClientComponentSection struct {
	ID   fields.FieldInt  `json:"id"`
	Name fields.FieldText `json:"name" required:"TRUE" max:"250"`
}

func (m ClientComponentSection) Relation() string {
	return clientComponentSectionRelation
}

func (m ClientComponentSection) CollectionAgg() any {
	return &TotCount{0}
}

type ClientComponentSectionKey struct {
	ID fields.FieldInt `json:"id" required:"true"`
}

func (m ClientComponentSectionKey) Relation() string {
	return clientComponentSectionRelation
}

type ClientComponentSectionUpdate struct {
	Keys  ClientComponentSectionKey `json:"keys"`
	Model ClientComponentSection    `json:"model"`
}

type ClientComponentSectionDelete struct {
	Keys []ClientComponentSectionKey `json:"keys"`
}
