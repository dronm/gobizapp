package models

import (
	fields "github.com/dronm/crudifier/fields"
)

const (
	mainMenuRelation = "main_menu"
)

type MainMenu struct {
	ID                fields.FieldInt  `json:"id" primaryKey:"true" srvCalc:"true"`
	Caption           fields.FieldText `json:"caption" required:"true" maxLen:"250"`
	RoleID            fields.FieldText `json:"role_id" enum:"role_id"`
	UserID            fields.FieldInt  `json:"user_id"`
	ClientComponentID fields.FieldInt  `json:"client_component_id"`
	ParentID          fields.FieldInt  `json:"parent_id"`
}

func (m MainMenu) Relation() string {
	return mainMenuRelation
}

func (m MainMenu) CollectionAgg() any {
	return &TotCount{0}
}

type MainMenuKey struct {
	ID fields.FieldInt `json:"id" required:"true"`
}

func (m MainMenuKey) Relation() string {
	return mainMenuRelation
}

type MainMenuUpdate struct {
	Keys  MainMenuKey `json:"keys"`
	Model MainMenu    `json:"model"`
}

type MainMenuDelete struct {
	Keys []MainMenuKey `json:"keys"`
}
