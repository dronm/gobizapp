package services

import (
	"context"
	"fmt"

	crud "github.com/dronm/crudifier"
	fields "github.com/dronm/crudifier/fields"
	crudTypes "github.com/dronm/crudifier/types"
	"github.com/dronm/ds/pgds"
	"github.com/dronm/session"

	"github.com/dronm/gobizapp/models"
)

// MainMenuService is a service for managing product categories
type MainMenuService struct {
	DB      *pgds.PgProvider
	Session session.Session
}

func NewMainMenuService(db *pgds.PgProvider, sess session.Session) *MainMenuService {
	return &MainMenuService{DB: db, Session: sess}
}

func (s *MainMenuService) FetchList(ctx context.Context, params crud.CollectionParams) ([]*models.MainMenuList, *models.TotCount, error) {
	return FetchCollectionModel(ctx, s.DB, &models.MainMenuList{}, &models.TotCount{}, params)
}

func (s *MainMenuService) FetchDetail(ctx context.Context, id int) (*models.MainMenuList, error) {
	model := models.MainMenuList{}
	if err := FetchModel(ctx, s.DB, &models.MainMenuKey{ID: fields.NewFieldInt(int64(id), true, false)}, &model); err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *MainMenuService) Delete(ctx context.Context, keyModels []models.MainMenuKey) (int64, error) {
	models := make([]crudTypes.DbModel, len(keyModels))
	for i, m := range keyModels {
		models[i] = m
	}
	cnt, err := DeleteModel(ctx, s.DB, models, nil)
	if err != nil {
		return 0, err
	}

	_ = EvHandler.PublishEvent(s.Session.SessionID(), "MainMenu.Delete", keyModels)

	return cnt, nil
}

func (s *MainMenuService) Update(ctx context.Context, keyModel models.MainMenuKey, model models.MainMenu) (int64, error) {
	cnt, err := UpdateModel(ctx, s.DB, keyModel, &model)
	if err != nil {
		return 0, err
	}

	_ = EvHandler.PublishEvent(s.Session.SessionID(), "MainMenu.Update", keyModel)

	return cnt, nil
}

func (s *MainMenuService) Insert(ctx context.Context, model models.MainMenu) (map[string]any, error) {
	retFields, err := InsertModel(ctx, s.DB, &model, nil)
	if err != nil {
		return nil, err
	}

	_ = EvHandler.PublishEvent(s.Session.SessionID(), "MainMenu.Insert", retFields)

	return retFields, nil
}

func (s *MainMenuService) MoveUp(ctx context.Context, itemID int) error {
	return s.MoveItem(ctx, itemID, "up")
}

func (s *MainMenuService) MoveDown(ctx context.Context, itemID int) error {
	return s.MoveItem(ctx, itemID, "down")
}

func (s *MainMenuService) MoveItem(ctx context.Context, itemID int, direction string) error {
	poolConn, connID, err := s.DB.GetPrimary()
	if err != nil {
		return fmt.Errorf("GetPrimary() failed: %v", err)
	}
	defer s.DB.Release(poolConn, connID)
	conn := poolConn.Conn()

	if _, err := conn.Exec(ctx,
		fmt.Sprintf("SELECT main_menu_item_%s($1)", direction),
		itemID,
	); err != nil {
		return fmt.Errorf("conn.Exec() main_menu_item_%s: %v", direction, err)
	}

	return nil
}
