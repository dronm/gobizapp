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

// ClientComponentSection is a service.
type ClientComponentSectionService struct {
	DB      *pgds.PgProvider
	Session session.Session
}

// NewClientComponentSectionService is a service constructor.
func NewClientComponentSectionService(db *pgds.PgProvider, sess session.Session) *ClientComponentSectionService {
	return &ClientComponentSectionService{DB: db, Session: sess}
}

// Insert inserts one row into models.ClientComponentSection
func (s *ClientComponentSectionService) Insert(ctx context.Context, model models.ClientComponentSection) (map[string]interface{}, error) {
	retFields, err := InsertModel(ctx, s.DB, &model, nil)
	if err != nil {
		return nil, err
	}
	return retFields, nil
}

func (s *ClientComponentSectionService) Update(ctx context.Context, keyModel models.ClientComponentSectionKey, model models.ClientComponentSection) (int64, error) {
	cnt, err := UpdateModel(ctx, s.DB, keyModel, &model)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// Delete removes row from models.ClientComponentSection
func (s *ClientComponentSectionService) Delete(ctx context.Context, keyModels []models.ClientComponentSectionKey) (int64, error) {
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

// FetchDetail retrieves one row from models.ClientComponentSection
func (s *ClientComponentSectionService) FetchDetail(ctx context.Context, id int) (*models.ClientComponentSection, error) {
	model := models.ClientComponentSection{}
	if err := FetchModel(ctx, s.DB, &models.ClientComponentSectionKey{ID: fields.NewFieldInt(int64(id), true, false)}, &model); err != nil {
		return nil, err
	}
	return &model, nil
}

// FetchList retrieves rows from models.ClientComponentSection. Returns two models: data and aggregates.
func (s *ClientComponentSectionService) FetchList(ctx context.Context, params crud.CollectionParams) ([]*models.ClientComponentSection, *models.TotCount, error) {
	return FetchCollectionModel(ctx, s.DB, &models.ClientComponentSection{}, &models.TotCount{}, params)
}
