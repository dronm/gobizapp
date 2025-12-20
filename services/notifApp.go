package services

import (
	"context"

	crud "github.com/dronm/crudifier"
	"github.com/dronm/ds/pgds"
	"github.com/dronm/session"

	"github.com/dronm/gobizapp/models"
)

type NotifAppConfiger interface {
	GetAppID() int
}
var notifAppConf NotifAppConfiger

// ProductCatService is a service for managing product categories
type NotifAppService struct {
	DB      *pgds.PgProvider
	Session session.Session
}

func NewNotifAppService(db *pgds.PgProvider, sess session.Session) *NotifAppService {
	return &NotifAppService{DB: db, Session: sess}
}

func (s *NotifAppService) FetchList(ctx context.Context, params crud.CollectionParams) ([]*models.NotifApp, *models.TotCount, error) {
	//add name param to filter
	params.Filter = append(params.Filter,
		crud.CollectionFilter{
			Join:   crud.FILTER_PAR_JOIN_AND,
			Fields: map[string]crud.CollectionFilterField{"id": {Operator: crud.FILTER_OPER_PAR_E, Value: NotifAppID()}},
		},
	)

	return FetchCollectionModel(ctx, s.DB, &models.NotifApp{}, &models.TotCount{}, params)
}

func (s *NotifAppService) FetchDetail(ctx context.Context) (*models.NotifApp, error) {
	model := models.NotifApp{}
	if err := FetchModel(ctx, s.DB, &models.NotifAppKey{Id: NotifAppID()}, &model); err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *NotifAppService) Update(ctx context.Context, model models.NotifAppNew) (int64, error) {
	keyModel := models.NotifAppKey{Id: NotifAppID()}
	cnt, err := UpdateModel(ctx, s.DB, keyModel, &model)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func NotifAppID() int {
	if notifAppConf != nil {
		return notifAppConf.GetAppID()
	}
	return 0
}
