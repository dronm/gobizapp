package models

const (
	notifAppsRelation = "apps"
)

type NotifTMParams struct {
	Token string `json:"token"`
}

type NotifWAParams struct {
	QRCode string `json:"qr_code"`
}

type NotifEmailParams struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

type NotifSMSParams struct {
	Sign   string `json:"sign"`
	Pwd    string `json:"pwd"`
	Login  string `json:"login"`
	Active bool   `json:"active"`
}

type NotifProviderParams struct {
	TM    NotifTMParams    `json:"tm"`
	Email NotifEmailParams `json:"email"`
	WA    NotifWAParams    `json:"wa"`
	SMS   NotifSMSParams   `json:"sms"`
}

// object model for insert/update
type NotifApp struct {
	Id             int                 `json:"id" required:"true" primaryKey:"true"`
	Name           string              `json:"name" required:"true"`
	ProviderParams NotifProviderParams `json:"provider_params" required:"true"`
	CallbackUrl    string              `json:"callback_url" required:"true"`
	CallbackKey    string              `json:"callback_key" required:"true"`
	Pwd            string              `json:"pwd" required:"true"`
}

func (m NotifApp) Relation() string {
	return notifAppsRelation
}

func (m NotifApp) CollectionAgg() any {
	return &TotCount{0}
}

// object key model
type NotifAppKey struct {
	Id int `json:"id" required:"true"`
}

func (m NotifAppKey) Relation() string {
	return notifAppsRelation
}

type NotifAppNew struct {
	ProviderParams NotifProviderParams `json:"provider_params" required:"true"`
	CallbackUrl    string              `json:"callback_url" required:"true"`
	CallbackKey    string              `json:"callback_key" required:"true"`
	Pwd            string              `json:"pwd" required:"true"`
}

func (m NotifAppNew) Relation() string {
	return notifAppsRelation
}

// update model
type NotifAppUpdate struct {
	Keys  NotifAppKey `json:"keys"`
	Model NotifAppNew `json:"model"`
}
