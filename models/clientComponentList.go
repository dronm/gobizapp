package models

const (
	clientComponentListRelation = "client_components_list"
)

type ClientComponentList struct {
	ID                         int     `json:"id"`
	Caption                    string  `json:"caption" alias:"Наименование"`
	ClientComponentSectionID   int     `json:"client_component_section_id"`
	ClientComponentSectionsRef Ref     `json:"client_component_sections_ref"`
	Name                       string  `json:"name" alias:"Vue name"`
	Path                       string  `json:"path"`
	Component                  string  `json:"component" alias:"Vue component"`
	CommentText                *string `json:"comment_text"`
}

func (m ClientComponentList) Relation() string {
	return clientComponentListRelation
}

func (m ClientComponentList) CollectionAgg() any {
	return &TotCount{0}
}
