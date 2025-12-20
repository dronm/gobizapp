package models

const (
	mainMenuListRelation = "main_menu_list"
)

// MainMenuList is a virtual table model.
type MainMenuList struct {
	ID                int     `json:"id" primaryKey:"true" srvCalc:"true"`
	Caption           string  `json:"caption" required:"true" maxLen:"250"`
	RoleID            *string `json:"role_id" enum:"role_id"`
	UserID            *int    `json:"user_id"`
	User              Ref     `json:"users_ref"`
	ClientComponentID *int    `json:"client_component_id"`
	ClientComponent   Ref     `json:"client_components_ref"`
	ParentID          *int    `json:"parent_id"`
}

func (m MainMenuList) Relation() string {
	return mainMenuListRelation
}

func (m MainMenuList) CollectionAgg() any {
	return &TotCount{0}
}
