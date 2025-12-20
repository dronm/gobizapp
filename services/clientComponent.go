package services

import (
	"context"
	"github.com/dronm/gobizapp/models"
	crud "github.com/dronm/crudifier"
	fields "github.com/dronm/crudifier/fields"
	crudTypes "github.com/dronm/crudifier/types"
	"github.com/dronm/ds/pgds"
	"github.com/dronm/session"
)

type ClientComponentService struct {
	DB      *pgds.PgProvider
	Session session.Session
}

// NewClientComponentService is a service constructor.
func NewClientComponentService(db *pgds.PgProvider, sess session.Session) *ClientComponentService {
	return &ClientComponentService{DB: db, Session: sess}
}

// Insert inserts one row into models.ClientComponent
func (s *ClientComponentService) Insert(ctx context.Context, model models.ClientComponent) (map[string]interface{}, error) {
	retFields, err := InsertModel(ctx, s.DB, &model, nil)
	if err != nil {
		return nil, err
	}
	return retFields, nil
}

func (s *ClientComponentService) Update(ctx context.Context, keyModel models.ClientComponentKey, model models.ClientComponent) (int64, error) {
	cnt, err := UpdateModel(ctx, s.DB, keyModel, &model)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// Delete removes row from models.ClientComponent
func (s *ClientComponentService) Delete(ctx context.Context, keyModels []models.ClientComponentKey) (int64, error) {
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

// FetchDetail retrieves one row from models.ClientComponent
func (s *ClientComponentService) FetchDetail(ctx context.Context, id int) (*models.ClientComponentList, error) {
	model := models.ClientComponentList{}
	if err := FetchModel(ctx, s.DB, &models.ClientComponentKey{ID: fields.NewFieldInt(int64(id), true, false)}, &model); err != nil {
		return nil, err
	}
	return &model, nil
}

// FetchList retrieves rows from models.ClientComponent. Returns two models: data and aggregates.
func (s *ClientComponentService) FetchList(ctx context.Context, params crud.CollectionParams) ([]*models.ClientComponentList, *models.TotCount, error) {
	return FetchCollectionModel(ctx, s.DB, &models.ClientComponentList{}, &models.TotCount{}, params)
}
