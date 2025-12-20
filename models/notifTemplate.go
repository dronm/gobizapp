package models

import (
	fields "github.com/dronm/crudifier/fields"
)

const (
	notifTemplateRelation      = "notif_templates"
	notifTemplateListRelation = "notif_templates_list"
)

type NotifTemplateField struct {
	Id string `json:"id"`
}

type NotifTemplateProviderValue struct {
	Id  string `json:"id"`
	Val string `json:"val"`
}

// object model for insert/update
type NotifTemplate struct {
	Id             fields.FieldInt               `json:"id" primaryKey:"true" srvCalc:"true"`
	NotifProvider  fields.FieldText              `json:"notif_provider" required:"true" enum:"notif_providers"`
	NotifType      fields.FieldText              `json:"notif_type" required:"true" enum:"notif_types"`
	Template       fields.FieldText              `json:"template" required:"true"`
	CommentText    fields.FieldText              `json:"comment_text"`
	Fields         []NotifTemplateField            `json:"fields"`
	ProviderValues *[]NotifTemplateProviderValue `json:"provider_values"`
}

func (m NotifTemplate) Relation() string {
	return notifTemplateRelation
}

func (m NotifTemplate) CollectionAgg() any {
	return &TotCount{0}
}

// object key model
type NotifTemplateKey struct {
	Id fields.FieldInt `json:"id" required:"true"`
}

func (m NotifTemplateKey) Relation() string {
	return notifTemplateRelation
}

// update model
type NotifTemplateUpdate struct {
	Keys  NotifTemplateKey `json:"keys"`
	Model NotifTemplate    `json:"model"`
}

// delete model
type NotifTemplateDelete struct {
	Keys []NotifTemplateKey `json:"keys"`
}

// ******************
type NotifTemplateList struct {
	Id            fields.FieldInt  `json:"id" primaryKey:"true" srvCalc:"true"`
	NotifProvider fields.FieldText `json:"notif_provider" required:"true" enum:"notif_providers"`
	NotifType     fields.FieldText `json:"notif_type" required:"true" enum:"notif_types"`
	Template      fields.FieldText `json:"template" required:"true"`
}

func (m NotifTemplateList) Relation() string {
	return notifTemplateListRelation
}
