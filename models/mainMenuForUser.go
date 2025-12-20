package models


const (
	mainMenuForUserRelation = "main_menu_for_user"
)


type MainMenuForUser struct {
	ID       int                `json:"id" primaryKey:"true" srvCalc:"true"`
	Caption  string             `json:"caption" required:"true" maxLen:"250"`
	Route    string             `json:"route" required:"true" maxLen:"250"`
	Children *[]MainMenuForUser `json:"children"`
}

func (m MainMenuForUser) Relation() string {
	return mainMenuForUserRelation
}

func (m MainMenuForUser) CollectionAgg() any {
	return &TotCount{0}
}

