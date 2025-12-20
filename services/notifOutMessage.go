package services

import (
	"context"

	crud "github.com/dronm/crudifier"
	"github.com/dronm/ds/pgds"
	"github.com/dronm/session"

	"github.com/dronm/gobizapp/models"
)

type NotifOutMessageService struct {
	DB      *pgds.PgProvider
	Session session.Session
}

func NewNotifOutMessageService(db *pgds.PgProvider, sess session.Session) *NotifOutMessageService {
	return &NotifOutMessageService{DB: db, Session: sess}
}

func (s *NotifOutMessageService) FetchList(ctx context.Context, params crud.CollectionParams) ([]*models.NotifOutMessage, *models.TotCount, error) {
	//add name param to filter
	params.Filter = append(params.Filter,
		crud.CollectionFilter{
			Join:   crud.FILTER_PAR_JOIN_AND,
			Fields: map[string]crud.CollectionFilterField{"app_id": {Operator: crud.FILTER_OPER_PAR_E, Value: NotifAppID()}},
		},
	)

	return FetchCollectionModel(ctx, s.DB, &models.NotifOutMessage{}, &models.TotCount{}, params)
}

func (s *NotifOutMessageService) FetchDetail(ctx context.Context, id int) (*models.NotifOutMessage, error) {
	model := models.NotifOutMessage{}
	if err := FetchModel(ctx, s.DB, &models.NotifOutMessageKey{Id: id, AppId: NotifAppID()}, &model); err != nil {
		return nil, err
	}
	return &model, nil
}

