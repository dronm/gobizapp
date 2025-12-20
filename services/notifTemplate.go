package services

import (
	"context"

	crud "github.com/dronm/crudifier"
	fields "github.com/dronm/crudifier/fields"
	crudTypes "github.com/dronm/crudifier/types"
	"github.com/dronm/ds/pgds"
	"github.com/dronm/session"

	"github.com/dronm/gobizapp/models"
)

// ProductCatService is a service for managing product categories
type NotifTemplateService struct {
	DB      *pgds.PgProvider
	Session session.Session
}

func NewNotifTemplateService(db *pgds.PgProvider, sess session.Session) *NotifTemplateService {
	return &NotifTemplateService{DB: db, Session: sess}
}

func (s *NotifTemplateService) FetchList(ctx context.Context, params crud.CollectionParams) ([]*models.NotifTemplate, *models.TotCount, error) {
	return FetchCollectionModel(ctx, s.DB, &models.NotifTemplate{}, &models.TotCount{}, params)
}

func (s *NotifTemplateService) FetchDetail(ctx context.Context, id int) (*models.NotifTemplate, error) {
	model := models.NotifTemplate{}
	if err := FetchModel(ctx, s.DB, &models.NotifTemplateKey{Id: fields.NewFieldInt(int64(id), true, false)}, &model); err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *NotifTemplateService) Delete(ctx context.Context, keyModels []models.NotifTemplateKey) (int64, error) {
	models := make([]crudTypes.DbModel, len(keyModels))
	for i, m := range keyModels {
		models[i] = m
	}
	cnt, err := DeleteModel(ctx, s.DB, models, nil)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (s *NotifTemplateService) Update(ctx context.Context, keyModel models.NotifTemplateKey, model models.NotifTemplate) (int64, error) {
	cnt, err := UpdateModel(ctx, s.DB, keyModel, &model)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (s *NotifTemplateService) Insert(ctx context.Context, model models.NotifTemplate) (map[string]interface{}, error) {
	retFields, err := InsertModel(ctx, s.DB, &model, nil)
	if err != nil {
		return nil, err
	}
	return retFields, nil
}

